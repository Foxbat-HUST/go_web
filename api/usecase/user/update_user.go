package user

import (
	"go_web/domain/entity"
	"go_web/domain/service"

	"gorm.io/gorm"
)

type UpdateUserInput struct {
	ID    uint32
	Name  string `json:"name"`
	Age   uint16 `json:"age"`
	Email string `json:"email"`
}

type UpdateUserOuput struct {
	entity.User
}

type UpdateUser interface {
	Exec(input UpdateUserInput) (ouput *UpdateUserOuput, err error)
}

type updateUser struct {
	db          *gorm.DB
	userService service.UserService
}

func NewUpdateUser(db *gorm.DB, userService service.UserService) UpdateUser {
	return &updateUser{
		db:          db,
		userService: userService,
	}
}

func (c *updateUser) Exec(input UpdateUserInput) (*UpdateUserOuput, error) {
	user := entity.User{
		ID:    input.ID,
		Name:  input.Name,
		Age:   input.Age,
		Email: input.Email,
	}

	if _, err := c.userService.GetByID(input.ID); err != nil {
		return nil, err
	}

	if err := c.userService.ValidateUpdate(user); err != nil {
		return nil, err
	}

	var updatedUser entity.User
	if err := c.db.Transaction(func(tx *gorm.DB) error {
		u, e := c.userService.WithTx(tx).Update(user.ID, user)
		if e != nil {
			return e
		}
		updatedUser = *u
		return nil
	}); err != nil {
		return nil, err
	}

	return &UpdateUserOuput{updatedUser}, nil
}
