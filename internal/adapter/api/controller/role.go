package api

import (
	"net/http"
	"online-script/internal/core/domain"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	roleUseCase domain.RoleUsecase
}

func NewRoleController(roleUseCase domain.RoleUsecase) *RoleController {
	return &RoleController{
		roleUseCase: roleUseCase,
	}
}

func (controller *RoleController) ListRoles(c *gin.Context) {

	var params domain.ListRolesParams
	err := c.ShouldBind(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	roles, err := controller.roleUseCase.ListRoles(c, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, roles)
}
