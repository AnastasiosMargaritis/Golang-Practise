package config

import (
	"github.com/jackc/pgx/v5"
)

type Application struct {
	Env      *Env
	Postgres *pgx.Conn
}

func App(path string) Application {
	app := &Application{}
	app.Env = NewEnv(path)
	app.Postgres = NewPostgresDB(app.Env)
	return *app
}
