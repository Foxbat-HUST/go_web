package auth

import (
	"go_web/domain/entity"
	"go_web/domain/service"
	"go_web/errors"
)

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginOutput struct {
	Token string
}

type Login interface {
	Exec(input LoginInput) (*LoginOutput, error)
}

type login struct {
	authService service.AuthService
	userService service.UserService
}

func NewLogin(authService service.AuthService, userService service.UserService) Login {
	return &login{
		authService: authService,
		userService: userService,
	}
}

func (l login) Exec(input LoginInput) (*LoginOutput, error) {
	var err error
	form := entity.LoginForm{
		Email:    input.Email,
		Password: input.Password,
	}

	err = l.authService.ValidateLoginForm(form)
	if err != nil {
		return nil, err
	}

	err = l.authService.AuthWithLoginForm(form)
	if err != nil {
		return nil, errors.Unauthorized(err)
	}

	var user *entity.User
	user, err = l.userService.GetByEmail(form.Email)
	if err != nil {
		return nil, err
	}

	var token string
	token, err = l.authService.CreateToken(*user)
	if err != nil {
		return nil, err
	}

	return &LoginOutput{Token: token}, nil
}
