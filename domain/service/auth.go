package service

import "go_web/domain/entity"

type AuthService interface {
	ValidateLoginForm(form entity.LoginForm) error
	AuthWithLoginForm(form entity.LoginForm) error
	HashPassword(rawPassword string) (string, error)
	CreateToken(user entity.User) (string, error)
	ParseToken(tokenStr string) (*entity.User, error)
	InvalidToken() error
}
