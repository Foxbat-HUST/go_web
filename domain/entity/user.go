package entity

import (
	"go_web/errors"
)

type User struct {
	ID    uint32   `json:"id"`
	Name  string   `json:"name" validate:"required"`
	Age   uint16   `json:"age" validate:"required"`
	Email string   `json:"email" validate:"required,email"`
	Type  UserType `json:"type" validate:"userType"`
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
