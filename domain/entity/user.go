package entity

import (
	"go_web/errors"
	"time"
)

type User struct {
	ID        uint32
	Name      string `validate:"required"`
	Age       uint16 `validate:"required"`
	Email     string `validate:"required,email"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) Valid() error {
	if err := validate.Struct(u); err != nil {
		return errors.BadRequest(err)
	}

	return nil
}
