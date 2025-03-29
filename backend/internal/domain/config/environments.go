package config

import (
	"os"

	"github.com/joho/godotenv"
)

// postgres configuration struct that holds the database connection information
type PGConfig struct {
	Username string
	Password string
	Host     string
	DBName   string
	Port     string
}

type SuperTokenConfig struct {
	ApiKey        string
	ConnectionUri string
}

type Config interface {
	GetPostgresConfig() PGConfig
	GetSuperTokensConfig() SuperTokenConfig
}

type config struct {
	pg                PGConfig
	superTokensConfig SuperTokenConfig
}

func NewConfig() (Config, error) {

	// load the configuration from .env file
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	pgConfig := PGConfig{
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Host:     os.Getenv("POSTGRES_HOST"),
		DBName:   os.Getenv("POSTGRES_DB"),
		Port:     os.Getenv("POSTGRES_PORT"),
	}

	superTokensConfig := SuperTokenConfig{
		ConnectionUri: os.Getenv("SUPERTOKENS_URI"),
		ApiKey:        os.Getenv("SUPERTOKENS_API_KEY"),
	}

	return &config{
		pg:                pgConfig,
		superTokensConfig: superTokensConfig,
	}, nil
}

func (c *config) GetPostgresConfig() PGConfig {
	return c.pg
}

func (c *config) GetSuperTokensConfig() SuperTokenConfig {
	return c.superTokensConfig
}
