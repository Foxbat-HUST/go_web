package user

import (
	"go_web/domain/entity"
	"go_web/domain/service"
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
	userService service.UserService
}

func NewGetUser(userService service.UserService) GetUser {
	return &getUser{
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
