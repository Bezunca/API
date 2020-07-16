package message_broker

import (
	"crypto/tls"
	"github.com/Bezunca/API/internal/config"
	"github.com/streadway/amqp"
	"log"
	"os"
	"time"
)

type Session struct {
	configs               *config.RabbitMQConfig
	tlsConfig             *tls.Config
	errorLog              *log.Logger
	connection            *amqp.Connection
	channel               *amqp.Channel
	done                  chan bool
	notifyConnectionClose chan *amqp.Error
	notifyChannelClose    chan *amqp.Error
	notifyConfirm         chan amqp.Confirmation
	isReady               bool
}

var globalSession *Session = nil

// New creates a new consumer state instance, and automatically
// attempts to connect to the server.
func New(
	configs *config.RabbitMQConfig,
	tlsConfig *tls.Config,
) (*Session, error) {

	globalSession = &Session{
		errorLog:  log.New(os.Stderr, "RabbitMQ ERROR - ", log.LUTC|log.Ldate|log.Lmsgprefix|log.Ltime),
		configs:   configs,
		tlsConfig: tlsConfig,
		done:      make(chan bool),
	}

	go globalSession.handleReconnect(configs.FormatRabbitMQURL())

	return globalSession, nil
}

// Get returns the current open session with RabbitMQ
func Get() *Session {
	if globalSession == nil {
		panic("RabbitMQ session must be initialized before used!")
	}
	return globalSession
}

// handleReconnect will wait for a connection error on
// notifyConnClose, and then continuously attempt to reconnect.
func (session *Session) handleReconnect(addr string) {
	for {
		if session.connection == nil || session.connection.IsClosed() {
			session.isReady = false
			log.Println("Attempting to connect")

			conn, err := session.connect(addr)
			if err != nil {
				session.errorLog.Printf("Failed to connect: Reason: %v. Retrying... ", err)

				select {
				case <-session.done:
					return
				case <-time.After(session.configs.ReconnectDelay):
				}
				continue
			}

			if done := session.handleReInit(conn); done {
				break
			}
		}
	}
}

// handleReconnect will wait for a channel error
// and then continuously attempt to re-initialize both channels
func (session *Session) handleReInit(conn *amqp.Connection) bool {
	for {
		session.isReady = false

		err := session.channelInit(conn)

		if err != nil {
			session.errorLog.Println("Failed to initialize channel. Retrying...")

			select {
			case <-session.done:
				return true
			case <-time.After(session.configs.ReconnectDelay):
			}
			continue
		}

		select {
		case <-session.done:
			return true
		case err := <-session.notifyConnectionClose:
			log.Printf("Connection closed (%s). Reconnecting...", err)
			return false
		case err := <-session.notifyChannelClose:
			if session.isReady {
				log.Printf("Channel closed (%s). Re-running channelInit...", err)
			} else {
				return true
			}
		}
	}
}

// channelInit will initialize channel & declare queue
func (session *Session) channelInit(conn *amqp.Connection) error {
	ch, err := conn.Channel()

	if err != nil {
		return err
	}

	err = ch.Confirm(false)

	if err != nil {
		return err
	}

	_, err = ch.QueueDeclare(
		session.configs.CEIQueue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	session.channel = ch
	session.notifyChannelClose = make(chan *amqp.Error)
	session.notifyConfirm = make(chan amqp.Confirmation, 1)
	session.channel.NotifyClose(session.notifyChannelClose)
	session.channel.NotifyPublish(session.notifyConfirm)

	session.isReady = true
	log.Println("Setup!")

	return nil
}

// connect will create a new AMQP connection
func (session *Session) connect(rabbitMQURL string) (*amqp.Connection, error) {
	conn, err := amqp.DialTLS(rabbitMQURL, session.tlsConfig)
	if err != nil {
		return nil, err
	}

	session.connection = conn
	session.notifyConnectionClose = make(chan *amqp.Error)
	session.connection.NotifyClose(session.notifyConnectionClose)
	session.isReady = true

	log.Println("Connected on RabbitMQ!")
	return conn, nil
}

// Close will cleanly shutdown the channel and connection.
func (session *Session) Close() error {
	if !session.isReady {
		return &AlreadyClosedError{}
	}
	session.isReady = false

	err := session.connection.Close()
	if err != nil {
		return err
	}
	close(session.done)
	return nil
}

// Push will push data onto the queue, and wait for a confirm.
// If no confirms are received until within the resendTimeout,
// it continuously re-sends messages until a confirm is received.
// This will block until the server sends a confirm. Errors are
// only returned if the push action itself fails
func (session *Session) Push(data []byte) error {
	if !session.isReady {
		return &NotConnectedError{}
	}
	for {
		err := session.channel.Publish(
			"",                       // Exchange
			session.configs.CEIQueue, // Routing key
			true,                     // Mandatory
			false,                    // Immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        data,
			},
		)
		if err != nil {
			session.errorLog.Println("Push failed. Retrying...")
			select {
			case <-session.done:
				return &ShutdownError{}
			case <-time.After(session.configs.ResendDelay):
			}
			continue
		}
		select {
		case confirm := <-session.notifyConfirm:
			if confirm.Ack {
				return nil
			} else {
				session.errorLog.Println("Nack received!")
			}
		case <-time.After(session.configs.ResendDelay):
		}
		session.errorLog.Println("Push didn't confirm. Retrying...")
	}
}

type (
	NotConnectedError  struct{}
	AlreadyClosedError struct{}
	ShutdownError      struct{}
)

func (e *NotConnectedError) Error() string {
	return "not connected to a server"
}

func (e *AlreadyClosedError) Error() string {
	return "already closed: not connected to the server"
}

func (e *ShutdownError) Error() string {
	return "session is shutting down"
}