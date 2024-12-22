package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"log"
)

var AppConfig *Config

type Config struct {
	ServerConfig ServerConfig
	DBConfig     DBConfig
	AppSecret    AppSecret
}

type ServerConfig struct {
	Port    string `env:"SERVER_PORT,required"`
	Version string `env:"SERVER_VERSION,required"`
}

type AppSecret struct {
	Secret string `env:"APP_SECRET"`
}

type DBConfig struct {
	DbHost     string `env:"DB_HOST,required"`
	DbPort     string `env:"DB_PORT,required"`
	DbUser     string `env:"DB_USER,required"`
	DbPassword string `env:"DB_PASSWORD,required"`
	DbName     string `env:"DB_NAME,required"`
	DbSslMode  string `env:"DB_SSLMODE,required"`
}

func LoadConfig() error {
	if err := godotenv.Load("app.env"); err != nil {
		log.Fatal("Error loading app.env file")
	}

	config := &Config{}

	if err := env.Parse(config); err != nil {
		log.Fatal("Error parsing config")
	}

	serverConfig := &ServerConfig{}

	if err := env.Parse(serverConfig); err != nil {
		log.Fatal("Error parsing config")
	}

	config.ServerConfig = *serverConfig

	dbConfig := &DBConfig{}
	if err := env.Parse(dbConfig); err != nil {
		log.Fatal("Error parsing config")
	}

	config.DBConfig = *dbConfig

	appSecret := &AppSecret{}
	if err := env.Parse(appSecret); err != nil {
		log.Fatal("Error parsing config")
	}

	config.AppSecret = *appSecret

	AppConfig = config

	return nil
}
