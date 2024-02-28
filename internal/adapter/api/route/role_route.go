package route

import (
	api "online-script/internal/adapter/api/controller"
	persistence "online-script/internal/adapter/outbound/peristence"
	"online-script/internal/core/usecase"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func NewRoleRouter(db *pgx.Conn, router *gin.Engine) {
	roleRepo := persistence.NewRoleRepository(db)
	roleUseCase := usecase.NewRoleUseCase(roleRepo)
	roleController := api.NewRoleController(roleUseCase)

	router.GET("/roles", roleController.ListRoles)
}
