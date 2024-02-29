package main

import (
	"online-script/internal/adapter/api/route"
	"online-script/internal/adapter/config"

	"github.com/gin-gonic/gin"
)

func main() {
	app := config.App("../")
	env := app.Env
	db := app.Postgres

	gin := gin.Default()

	route.Setup(db, gin)
	gin.Run(env.ServerAddress)
}
