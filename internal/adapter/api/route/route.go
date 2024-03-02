package route

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func Setup(db *pgx.Conn, gin *gin.Engine) {
	// Register routes
	publicRouter := gin.Group("")
	NewRoleRouter(db, publicRouter)
	NewUserRouter(db, publicRouter)
}
