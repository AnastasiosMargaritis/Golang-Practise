package route

import (
	api "online-script/internal/adapter/api/controller"
	persistence "online-script/internal/adapter/outbound/peristence"
	"online-script/internal/core/usecase"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func NewUserRouter(db *pgx.Conn, group *gin.RouterGroup) {
	userRepository := persistence.NewUserRepository(db)
	userController := &api.UserController{
		UserUseCase: usecase.NewUserUseCase(userRepository),
	}
	group.POST("/createUser", userController.CreateUser)
}
