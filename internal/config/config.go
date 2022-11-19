package config

import (
	"errors"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	SECRET              string
	JWT_EXPIRATION_TIME string
	DB_HOST             string
	DB_PORT             string
	DB_PORT2            string
	DB_KEYSPACE         string
	RPC_PORT            string
	REST_PORT           string
	BROKER_PORT         string
	BROKER_HOST         string
}

func GetCfg() (Config, error) {
	cfg := Config{}
	err := readEnvironmentVariables(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func readEnvironmentVariables(cfg *Config) (err error) {
	err = godotenv.Load(".env")
	if err != nil {
		return err
	}
	vars := []string{
		"SECRET",
		"JWT_EXPIRATION_TIME",
		"DB_HOST",
		"DB_PORT",
		"DB_PORT2",
		"DB_KEYSPACE",
		"RPC_PORT",
		"REST_PORT",
		"BROKER_HOST",
		"BROKER_PORT"}

	for i := range vars {
		if os.Getenv(vars[i]) == "" {
			return errors.New("empty environment variable " + vars[i])
		}
	}
	cfg.DB_PORT = os.Getenv("DB_PORT")
	cfg.BROKER_HOST = os.Getenv("BROKER_HOST")
	cfg.BROKER_PORT = os.Getenv("BROKER_PORT")
	cfg.DB_PORT2 = os.Getenv("DB_PORT2")
	cfg.DB_HOST = os.Getenv("DB_HOST")
	cfg.DB_KEYSPACE = os.Getenv("DB_KEYSPACE")
	cfg.SECRET = os.Getenv("SECRET")
	cfg.RPC_PORT = os.Getenv("RPC_PORT")
	cfg.REST_PORT = os.Getenv("REST_PORT")
	cfg.JWT_EXPIRATION_TIME = os.Getenv("JWT_EXPIRATION_TIME")
	_, err = time.ParseDuration(cfg.JWT_EXPIRATION_TIME)
	if err != nil {
		return err
	}
	return nil
}
