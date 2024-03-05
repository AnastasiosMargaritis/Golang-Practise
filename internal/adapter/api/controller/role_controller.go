package api

import (
	"net/http"
	"online-script/internal/core/domain"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	RoleUseCase domain.RoleUsecase
}

func (controller *RoleController) ListRoles(c *gin.Context) {
	roles, err := controller.RoleUseCase.ListRoles(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, roles)
}
