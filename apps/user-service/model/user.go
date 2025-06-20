package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID  `db:"id"`
	Name         string     `db:"name"`
	Email        string     `db:"email"`
	PasswordHash string     `db:"password_hash"`
	CreatedAt    time.Time  `db:"created_at"`
	UpdatedAt    time.Time  `db:"updated_at"`
	DeletedAt    *time.Time `db:"deleted_at"`
}
