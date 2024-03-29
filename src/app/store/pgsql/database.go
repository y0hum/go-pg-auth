package pgsql

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type (
	DatabaseConfig struct {
		DBHost         string `yaml:"db_host"`
		DBPort         string `yaml:"db_port"`
		DBUsername     string `yaml:"db_username"`
		DBPassword     string `yaml:"db_password"`
		Database       string `yaml:"database"`
		UsersTable     string `yaml:"users_table"`
		LoginColumn    string `yaml:"login_column"`
		PasswordColumn string `yaml:"password_column"`
	}
)

func NewPoolInstance() *pgxpool.Pool {
	configFile, err := ioutil.ReadFile("config/database.yml")
	if err != nil {
		log.Fatal("Cant read database config", zap.String("details", err.Error()))
	}

	var config DatabaseConfig
	if err := yaml.Unmarshal(configFile, &config); err != nil {
		log.Fatal("Cant unmarshal database config", zap.String("details", err.Error()))
	}

	databaseURL := "postgresql://" + config.DBUsername + ":" + config.DBPassword + "@" + config.DBHost + ":" + config.DBPort + "/" + config.Database

	pool, err := pgxpool.Connect(context.Background(), databaseURL)
	if err != nil {
		log.Fatal("Unable to connect to PostgreSQL database", zap.String("details", err.Error()))
	}

	return pool
}
