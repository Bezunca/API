package config

import (
	"fmt"
	"github.com/Bezunca/mongo_connection/config"
	"github.com/fogodev/openvvar"
	"github.com/labstack/echo/v4"
	"strings"
	"time"
)

type (
	// Let's Encrypt Stuff
	acme struct {
		UseTLS         bool     `config:"use-tls;default=true;description=Whether to use TLS or not"`
		Environment    string   `config:"environment;required;options=STAGING, PROD;description=Let's Encrypt environment (STAGING or PROD)."`
		Domains        []string `config:"domains;required"`
		CacheDirectory string   `config:"cache-directory;required;description=Directory to save certificates from Let's Encrypt"`
		Email          string   `config:"email;default=bezuncainvestimentos@gmail.com.br"`
	}

	// JWT Stuff
	jwt struct {
		SecretAuth  string `config:"secret-auth;required"`
		SecretEmail string `config:"secret-email;required"`
	}

	database struct {
		Name           string `config:"name;default=bezunca"`
		UserCollection string `config:"user-collection;default=users"`
	}

	RabbitMQConfig struct {
		CEIQueue       string        `config:"cei-queue;required"`
		User           string        `config:"user;required"`
		Password       string        `config:"password;required"`
		Host           string        `config:"host;required"`
		AMQPPort       int           `config:"amqpport;required"`
		VHost          string        `config:"vhost;"`
		ReconnectDelay time.Duration `config:"reconnect-delay;default=2s"`
		RestartDelay   time.Duration `config:"restart-delay;default=1s"`
		ResendDelay    time.Duration `config:"resend-delay;default=1s"`
	}

	Config struct {
		Debug   bool   `config:"debug;default=false"`
		Address string `config:"address;default=localhost"`
		Port    string `config:"port;default=8080"`

		SendGridAPIKEY      string `config:"sendgrid-api-key;default=key"`
		WebURL              string `config:"web-url;default=url"`
		DynamicLink         string `config:"dynamic-link;default=dynamic"`
		FlutterAndroidAppID string `config:"flutter-android-app-id;default=id"`
		CAFilePath          string `config:"ca-file-path;required"`
		RSAPublicKeyPath    string `config:"rsa-public-key-path;required"`

		// B3 Fetch Prices Stuff
		B3CacheTimeout time.Duration `config:"b3-cache-timeout;default=10m"`

		JWT      jwt
		MongoDB  config.MongoConfigs
		ACME     acme
		Database database
		RabbitMQ RabbitMQConfig
	}
)

func (c *Config) ApplicationAddress() string {
	return strings.Join([]string{c.Address, c.Port}, ":")
}

func (c *RabbitMQConfig) FormatRabbitMQURL() string {
	return fmt.Sprintf("amqps://%s:%s@%s:%d/%s", c.User, c.Password, c.Host, c.AMQPPort, c.VHost)
}

var globalConfig *Config = nil

func New(log echo.Logger, notifyCallbacks ...func()) *Config {
	if globalConfig == nil {
		globalConfig = new(Config)
		if err := openvvar.Load(globalConfig); err != nil {
			log.Fatalf("An error occurred for bad config reasons: %v", err)
		}
	}

	// In this way we can notify the functions using the configs that the configs are ready
	for _, callback := range notifyCallbacks {
		callback()
	}

	return globalConfig
}

func Get() *Config {
	if globalConfig == nil {
		panic("Trying to get a nil config, you must use New function to instantiate configs before getting it")
	}
	return globalConfig
}
