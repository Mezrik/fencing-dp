package user

import (
	"github.com/google/uuid"
)

type UserRepository interface {
	GetUser(id uuid.UUID) (*User, error)
}
