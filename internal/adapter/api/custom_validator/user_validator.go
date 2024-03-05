package custom_validator

import (
	"context"
	"online-script/internal/core/usecase"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"

	persistence "online-script/internal/adapter/outbound/peristence"
)

func NewUserValidator(db *pgx.Conn) {
	userRepository := persistence.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepository)

	// Custom validation function
	var validUsername validator.Func = func(fl validator.FieldLevel) bool {
		// Retrieve the username value
		username := fl.Field().String()

		user, _ := userUseCase.GetByUsername(context.Background(), username)
		return user.Username == ""
	}

	if validator, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validator.RegisterValidation("validateUsername", validUsername)
	}
}
