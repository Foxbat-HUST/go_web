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

type CreateUserOuput struct {
	entity.User
}

type CreateUser interface {
	Exec(input CreateUserInput) (ouput *CreateUserOuput, err error)
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

func (c *createUser) Exec(input CreateUserInput) (*CreateUserOuput, error) {
	user := entity.User{
		Name:  input.Name,
		Age:   input.Age,
		Email: input.Email,
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

	return &CreateUserOuput{createdUser}, nil
}
