package api

import (
	"net/http"
	"online-script/internal/core/domain"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUseCase domain.UserUseCase
}

func (controller *UserController) CreateUser(c *gin.Context) {
	var req domain.CreateUserReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
	}

	createdUser, err := controller.UserUseCase.CreateUser(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}
