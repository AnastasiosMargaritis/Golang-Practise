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

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	createdUser, err := controller.UserUseCase.CreateUser(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Username in use."})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}
