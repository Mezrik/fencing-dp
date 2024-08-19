package user

import (
	"time"

	"github.com/Mezrik/fencing-dp/internal/common"
	"github.com/google/uuid"
)

type UserRoleEnum string

const (
	Admin     UserRoleEnum = "admin"
	BasicUser UserRoleEnum = "basic_user"
	Guest     UserRoleEnum = "guest"
	Unknow    UserRoleEnum = "unknown"
)

type User struct {
	common.Entity
	displayName string
	email       string
	role        UserRoleEnum
}

func NewUser(displayName string, email string, role UserRoleEnum) *User {
	return &User{
		Entity:      common.Entity{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: nil},
		displayName: displayName,
		email:       email,
		role:        role,
	}
}

func (user *User) DisplayName() string {
	return user.displayName
}

func (user *User) Email() string {
	return user.email
}

func (user *User) Role() UserRoleEnum {
	return user.role
}
