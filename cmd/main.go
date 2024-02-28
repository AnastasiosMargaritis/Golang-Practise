package main

import (
	"context"
	"fmt"
	"log"
	"online-script/internal/adapter/api/route"
	"online-script/internal/adapter/config"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func main() {
	gin := gin.Default()

	conf, err := config.LoadConfig("../")
	if err != nil {
		log.Fatal("Configuration can't be loaded: ", err)
	}

	var dbURL string = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", conf.DBUser, conf.DBPassword, conf.DBHost, conf.DBPort, conf.DBName)

	conn, err := pgx.Connect(context.Background(), dbURL)

	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	route.Setup(conn, gin)

	gin.Run(conf.ServerAddress)
}
