package config

import (
	"github.com/Bezunca/mongo_connection/config"
	"github.com/fogodev/openvvar"
	"github.com/labstack/echo/v4"
	"strings"
)

type Config struct {
	Debug               bool   `config:"debug;default=false"`
	Address             string `config:"address;default=localhost"`
	Port                string `config:"port;default=8080"`
	JWTSecretAuth       string `config:"jwt-secret-auth;default=secret"`
	JWTSecretEmail      string `config:"jwt-secret-email;default=secret"`
	SendGridAPIKEY      string `config:"sendgrid-api-key;default=key"`
	WebURL              string `config:"web-url;default=url"`
	DynamicLink         string `config:"dynamic-link;default=dynamic"`
	FlutterAndroidAppID string `config:"flutter-android-app-id;default=id"`
	MongoDB             config.MongoConfigs
	CAFile              string `config:"ca-file;required"`
	RSAPublicKey        string `config:"rsa-public-key;required"`
}

func (c *Config) ApplicationAddress() string {
	return strings.Join([]string{c.Address, c.Port}, ":")
}

var globalConfig *Config = nil

func New(log echo.Logger) *Config {
	if globalConfig == nil {
		globalConfig = new(Config)
		if err := openvvar.Load(globalConfig); err != nil {
			log.Fatalf("An error occurred for bad config reasons: %v", err)
		}
	}

	return globalConfig
}

func Get() *Config {
	if globalConfig == nil {
		panic("Trying to get a nil config, you must use New function to instantiate configs before getting it")
	}
	return globalConfig
}
