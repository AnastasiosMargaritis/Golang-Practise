package route

import (
	api "online-script/internal/adapter/api/controller"
	persistence "online-script/internal/adapter/outbound/peristence"
	"online-script/internal/core/usecase"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func NewRoleRouter(db *pgx.Conn, group *gin.RouterGroup) {
	roleRepo := persistence.NewRoleRepository(db)
	roleController := &api.RoleController{
		RoleUseCase: usecase.NewRoleUseCase(roleRepo),
	}

	group.GET("/roles/getAllRoles", roleController.ListRoles)
}
