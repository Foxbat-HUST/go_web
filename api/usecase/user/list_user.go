package user

import (
	"go_web/domain/entity"
	"go_web/domain/service"
)

type ListUserInput struct {
	PageIndex   int
	ItemPerPage int
}

type ListUserOutput struct {
	Users []*entity.User `json:"users"`
	Count int64          `json:"count"`
}

type ListUser interface {
	Exec(ListUserInput) (*ListUserOutput, error)
}

type listUserImpl struct {
	userService service.UserService
}

func NewListUser(userService service.UserService) ListUser {
	return &listUserImpl{
		userService: userService,
	}
}

func (l *listUserImpl) Exec(input ListUserInput) (*ListUserOutput, error) {
	data, cnt, err := l.userService.GetList(service.GetListOption{
		PageIndex:   input.PageIndex,
		ItemPerPage: input.ItemPerPage,
	})

	if err != nil {
		return nil, err
	}

	return &ListUserOutput{
		Users: data,
		Count: cnt,
	}, nil
}
