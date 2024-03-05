package route

import (
	api "online-script/internal/adapter/api/controller"
	"online-script/internal/adapter/config"
	persistence "online-script/internal/adapter/outbound/peristence"
	"online-script/internal/core/usecase"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func NewLoginRouter(env *config.Env, db *pgx.Conn, group *gin.RouterGroup) {
	userRepository := persistence.NewUserRepository(db)
	loginControlelr := &api.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(userRepository),
		Env:          env,
	}
	group.POST("/login", loginControlelr.Login)
}
