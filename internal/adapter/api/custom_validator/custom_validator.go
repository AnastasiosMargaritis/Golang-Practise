package custom_validator

import "github.com/jackc/pgx/v5"

func Setup(db *pgx.Conn) {
	NewUserValidator(db)
}
