package user

import (
	"go_web/domain/entity"
	"go_web/domain/service"

	"gorm.io/gorm"
)

type CreateUserInput struct {
	Name  string `json:"name"`
	Age   uint16 `json:"age"`
	Email string `json:"email"`
}

type CreateUserOutput struct {
	entity.User
}

type CreateUser interface {
	Exec(input CreateUserInput) (output *CreateUserOutput, err error)
}

type createUser struct {
	db          *gorm.DB
	userService service.UserService
}

func NewCreateUser(db *gorm.DB, userService service.UserService) CreateUser {
	return &createUser{
		db:          db,
		userService: userService,
	}
}

func (c *createUser) Exec(input CreateUserInput) (*CreateUserOutput, error) {
	user := entity.User{
		Name:  input.Name,
		Age:   input.Age,
		Email: input.Email,
		Type:  entity.UserTypeNormal,
	}
	if err := c.userService.ValidateCreate(user); err != nil {
		return nil, err
	}

	var createdUser entity.User
	if err := c.db.Transaction(func(tx *gorm.DB) error {
		u, e := c.userService.WithTx(tx).Create(user)
		if e != nil {
			return e
		}
		createdUser = *u
		return nil
	}); err != nil {
		return nil, err
	}

	return &CreateUserOutput{createdUser}, nil
}
