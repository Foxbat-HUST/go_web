package auth

import (
	"go_web/domain/entity"
	"go_web/domain/service"
	"go_web/errors"
)

type AuthInitInput struct {
	Token string
}

type AuthInitOutput struct {
	*entity.User `json:"user"`
	Token        string `json:"token"`
}

type AuthInit interface {
	Exec(AuthInitInput) (*AuthInitOutput, error)
}

type authInitImpl struct {
	authService service.AuthService
}

func NewAuthInit(
	authService service.AuthService,
) AuthInit {
	return &authInitImpl{
		authService: authService,
	}
}

func (a *authInitImpl) Exec(input AuthInitInput) (*AuthInitOutput, error) {
	user, err := a.authService.ParseToken(input.Token)
	if err != nil {
		return nil, errors.Unauthorized(err)
	}

	return &AuthInitOutput{
		user,
		input.Token,
	}, nil
}
