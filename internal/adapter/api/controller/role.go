package api

import (
	"net/http"
	"online-script/internal/core/domain"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	RoleUseCase domain.RoleUsecase
}

type ListRolesRequest struct {
	PageID   int32 `form:"page" binding:"required,min=1"`
	PageSize int32 `form:"size" binding:"required,min=5,max=10"`
}

func (controller *RoleController) ListRoles(c *gin.Context) {
	var req ListRolesRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	arg := domain.ListRolesParams{
		Limit:  req.PageID,
		Offset: req.PageSize,
	}

	roles, err := controller.RoleUseCase.ListRoles(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, roles)
}
