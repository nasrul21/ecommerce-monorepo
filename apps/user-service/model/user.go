package model

import (
	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `db:"id"`
	Name         string    `db:"name"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
}
