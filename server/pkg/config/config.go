package config

import (
	"fmt"
	"os"
	"stephen/server/pkg/db"
	"database/sql"
)

type Config struct {
	Host     string `yaml:"DbHost"`
	DBName     string `yaml:"DbName"`
	Password     string `yaml:"DbPass"`
	Username   string `yaml:"DbUser"`
	Port string `yaml:"ServerPort"`
	DB *sql.DB
}

var (
	host     = "127.0.0.1"
	dbName     = "certs"
	password     = "Password"
	username     = "root"
	port = "3306"
)

func NewConfig() (*Config, error) {
	var err error
	cfg := &Config{
		Host:     host,
		DBName:     dbName,
		Password:     password,
		Username:     username,
		Port: port,
	}

	// update config values from env, if any
	cfg.GETENVs()

	create := fmt.Sprintf("%s:%s@tcp(%s:%s)/", cfg.Username, cfg.Password, cfg.Host, cfg.Port )
	cfg.DB = db.InitDB(create)

	return cfg, err
}

func (c *Config) GETENVs() {
	if val, found := os.LookupEnv("DB_HOST"); found {
		c.Host = val
	}

	if val, found := os.LookupEnv("DB_NAME"); found {
		c.DBName = val
	}

	if val, found := os.LookupEnv("DB_PASS"); found {
		c.Password = val
	}

	if val, found := os.LookupEnv("DB_USER"); found {
		c.Username = val
	}

	if val, found := os.LookupEnv("DB_PORT"); found {
		c.Port = val
	}
}
