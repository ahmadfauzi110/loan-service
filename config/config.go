package config

import (
	"log"

	"github.com/spf13/viper"
)

var CurrentConfig *Config

type (
	Config struct {
		// DB configuration
		DB          *DB
		BASE_URL    string
		STATIC_PATH string
		BREVO       *BREVO
	}

	DB struct {
		Host     string
		Port     int
		Username string
		Password string
		Database string
		Debug    string
		SSLMode  string
	}

	BREVO struct {
		API_KEY      string
		SENDER_EMAIL string
		SENDER_NAME  string
	}
)

func Initialize(path string) *Config {
	var mappedCfg MappedConfig

	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	if err := viper.Unmarshal(&mappedCfg); err != nil {
		log.Fatal(err)
	}

	config := &Config{
		DB: &DB{
			Host:     mappedCfg.DBHost,
			Port:     mappedCfg.DBPort,
			Username: mappedCfg.DBUser,
			Password: mappedCfg.DBPassword,
			Database: mappedCfg.DBName,
			Debug:    mappedCfg.DBDebug,
			SSLMode:  mappedCfg.DBSSLMode,
		},
		BASE_URL:    mappedCfg.BASEURL,
		STATIC_PATH: mappedCfg.STATICPATH,
		BREVO: &BREVO{
			API_KEY:      mappedCfg.BREVOAPIKey,
			SENDER_EMAIL: mappedCfg.BREVOSenderEmail,
			SENDER_NAME:  mappedCfg.BREVOSenderName,
		},
	}

	CurrentConfig = config

	return config
}

type MappedConfig struct {
	// DB configuration
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     int    `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	DBDebug    string `mapstructure:"DB_DEBUG"`
	DBSSLMode  string `mapstructure:"DB_SSL_MODE"`

	BASEURL    string `mapstructure:"BASE_URL"`
	STATICPATH string `mapstructure:"STATIC_PATH"`

	BREVOAPIKey      string `mapstructure:"BREVO_API_KEY"`
	BREVOSenderEmail string `mapstructure:"BREVO_SENDER_EMAIL"`
	BREVOSenderName  string `mapstructure:"BREVO_SENDER_NAME"`
}
