package config

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

type Config struct {
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBName        string `mapstructure:"DB_NAME"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func NewPostgresDB(env *Env) *pgx.Conn {
	dbHost := env.DBHost
	dbPort := env.DBPort
	dbUser := env.DBUser
	dbPass := env.DBPassword
	dbName := env.DBName

	var dbURL string = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	conn, err := pgx.Connect(context.Background(), dbURL)

	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	return conn
}
