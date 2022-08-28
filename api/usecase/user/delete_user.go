package user

import (
	"go_web/domain/service"

	"gorm.io/gorm"
)

type DeleteUserInput struct {
	ID uint32
}

type DeleteUserOuput struct {
}

type DeleteUser interface {
	Exec(input DeleteUserInput) (ouput *DeleteUserOuput, err error)
}

type deleteUser struct {
	db          *gorm.DB
	userService service.UserService
}

func NewDeleteUser(db *gorm.DB, userService service.UserService) DeleteUser {
	return &deleteUser{
		db:          db,
		userService: userService,
	}
}

func (c *deleteUser) Exec(input DeleteUserInput) (*DeleteUserOuput, error) {

	if _, err := c.userService.GetByID(input.ID); err != nil {
		return nil, err
	}

	if err := c.db.Transaction(func(tx *gorm.DB) error {
		if e := c.userService.WithTx(tx).DeleteByID(input.ID); e != nil {
			return e
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &DeleteUserOuput{}, nil
}
