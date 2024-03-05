package route

import (
	"online-script/internal/adapter/api/middleware"
	"online-script/internal/adapter/config"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func Setup(env *config.Env, db *pgx.Conn, gin *gin.Engine) {

	// Public Routes
	publicRouter := gin.Group("")
	NewLoginRouter(env, db, publicRouter)
	NewUserRouter(db, publicRouter)

	// Register routes
	protectedRouter := gin.Group("")

	// Middleware to verify access
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// Private API's
	NewRoleRouter(db, protectedRouter)
}
