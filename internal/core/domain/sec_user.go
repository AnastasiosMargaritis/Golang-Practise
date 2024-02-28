package domain

import "github.com/jackc/pgtype"

type SecUser struct {
	Username       string
	HashedPassword string
	FirstName      string
	LastName       string
	Email          string
	CreatedAt      pgtype.Timestamptz
	RoleID         int64
}
