package entity

import (
	"go_web/errors"
	"time"
)

type User struct {
	ID        uint32
	Name      string   `validate:"required"`
	Age       uint16   `validate:"required"`
	Email     string   `validate:"required,email"`
	Type      UserType `validate:"userType"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserType string

const (
	UserTypeSuper  UserType = "super"
	UserTypeAdmin  UserType = "admin"
	UserTypeNormal UserType = "normal"
)

func (u UserType) IsValid() bool {
	if len(u) == 0 {
		return true
	}
	return u == UserTypeSuper || u == UserTypeAdmin || u == UserTypeNormal
}

func (u *User) Validate() error {
	if err := validate.Struct(u); err != nil {
		return errors.BadRequest(err)
	}

	return nil
}
