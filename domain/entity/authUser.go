package entity

import "go_web/errors"

type AuthUser struct {
	ID       uint32
	Email    string
	Type     UserType
	Password string
}

type LoginForm struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

func (l LoginForm) Validate() error {
	if err := validate.Struct(l); err != nil {
		return errors.BadRequest(err)
	}

	return nil
}
