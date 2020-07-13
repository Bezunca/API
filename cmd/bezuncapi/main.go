package main

import (
	"bezuncapi/internal/app"
	"bezuncapi/internal/app/controllers/b3"
	"bezuncapi/internal/app/middleware"
	"bezuncapi/internal/config"
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"golang.org/x/crypto/acme/autocert"

	"github.com/Bezunca/mongo_connection"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/acme"
)

func main() {
	// Echo instance
	e := echo.New()

	// Application configs
	configs := config.New(e.Logger, b3.SetFetcher)

	// Certificate Authority Stuff
	caFileBytes, err := ioutil.ReadFile(configs.CAFilePath)
	if err != nil {
		e.Logger.Fatal(err)
	}
	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM(caFileBytes)
	if !ok {
		e.Logger.Fatal("unable to parse CA Chain file")
	}
	tlsConfig := &tls.Config{
		RootCAs: roots,
	}

	// TLS Setup
	var startFunction func(string) error
	if configs.ACME.UseTLS {

		tlsCacheDir := filepath.Join(configs.ACME.CacheDirectory, "production")
		e.AutoTLSManager.Email = configs.ACME.Email

		if configs.ACME.Environment == "STAGING" {
			e.AutoTLSManager.Client = &acme.Client{
				DirectoryURL: "https://acme-staging-v02.api.letsencrypt.org/directory",
			}
			tlsCacheDir = filepath.Join(configs.ACME.CacheDirectory, "staging")
		}

		e.AutoTLSManager.HostPolicy = autocert.HostWhitelist(configs.ACME.Domains...)
		e.AutoTLSManager.Cache = autocert.DirCache(tlsCacheDir)

		startFunction = e.StartAutoTLS
	} else {
		startFunction = e.Start
	}

	// Initialize MongoDB
	if _, err := mongo_connection.New(&configs.MongoDB, tlsConfig); err != nil {
		e.Logger.Fatal(err)
	}

	// Pre Middleware
	middleware.PreMiddleware(e)

	// Middleware
	middleware.Middleware(e)

	// Routes
	app.Routes(e)

	if configs.Debug {
		e.Logger.SetLevel(log.DEBUG)
	}

	// Start server
	go func() {
		if err := startFunction(configs.ApplicationAddress()); err != nil {
			e.Logger.Errorf("shutting down the server and had some errors: %v", err)
		}
	}()

	waitShutdown(e)
}

func shutdownStuff(logger echo.Logger) {
	if err := mongo_connection.Close(); err != nil {
		logger.Error(err)
	}

}

func waitShutdown(server *echo.Echo) {
	// Wait for interrupt signal to gracefully waitShutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	shutdownStuff(server.Logger)

	server.Logger.Warn("Docking Bay self destruction sequence engaged!")
	if err := server.Shutdown(ctx); err != nil {
		server.Logger.Fatal(err)
	}
}
