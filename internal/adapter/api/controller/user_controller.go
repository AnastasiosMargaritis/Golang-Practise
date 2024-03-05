package api

import (
	"net/http"
	"online-script/internal/core/domain"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	UserUseCase domain.UserUseCase
}

func (controller *UserController) CreateUser(c *gin.Context) {
	var req domain.CreateUserReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Bind JSON Error"})
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Password Encryption Error"})
		return
	}

	req.Password = string(encryptedPassword)

	createdUser, err := controller.UserUseCase.CreateUser(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "User Creation Error"})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}
