package user

import (
	"go_web/domain/entity"
	"go_web/domain/service"

	"gorm.io/gorm"
)

type GetUserInput struct {
	ID uint32
}

type GetUserOutput struct {
	entity.User
}

type GetUser interface {
	Exec(input GetUserInput) (output *GetUserOutput, err error)
}

type getUser struct {
	db          *gorm.DB
	userService service.UserService
}

func NewGetUser(db *gorm.DB, userService service.UserService) GetUser {
	return &getUser{
		db:          db,
		userService: userService,
	}
}

func (c *getUser) Exec(input GetUserInput) (*GetUserOutput, error) {
	user, err := c.userService.GetByID(input.ID)
	if err != nil {
		return nil, err
	}

	return &GetUserOutput{*user}, nil
}
