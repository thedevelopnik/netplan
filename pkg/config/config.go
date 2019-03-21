package config

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type DBConfig struct {
	Host     string
	Port     string
	SSLMode  string
	User     string
	Password string
	DBName   string
}

type Config interface {
	GetDBConfig() DBConfig
}

func New() Config {
	v := viper.New()
	if strings.Compare(os.Getenv("ENV"), "production") != 0 {
		v.SetConfigName("config")
		v.AddConfigPath(".")
		v.SetConfigType("yaml")
		err := v.ReadInConfig()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		v.AutomaticEnv()
	}

	dbConfig := DBConfig{
		Host:     v.GetString("pg_host"),
		Port:     v.GetString("pg_port"),
		SSLMode:  v.GetString("pg_sslmode"),
		User:     v.GetString("pg_user"),
		Password: v.GetString("pg_password"),
		DBName:   v.GetString("pg_dbname"),
	}

	return config{
		DB: dbConfig,
	}
}

type config struct {
	DB DBConfig
}

func (c config) GetDBConfig() DBConfig {
	return c.DB
}
