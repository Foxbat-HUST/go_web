//go:build wireinject
// +build wireinject

package handler

import (
	"github.com/gin-gonic/gin"
	"go_web/api/usecase/auth"
	"go_web/api/usecase/user"
	"go_web/config"
	"go_web/domain/repository"
	"go_web/domain/service"
	"go_web/domain/service/implement"
	"go_web/infra/repository/mysql"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func initUserRepository(db *gorm.DB) repository.UserRepo {
	wire.Build(
		mysql.NewUserRepo,
	)

	return nil
}

func initUserService(db *gorm.DB) service.UserService {
	wire.Build(
		mysql.NewUserRepo,
		implement.NewUserService,
	)
	return nil
}

func InitAuthService(db *gorm.DB, cfg *config.Config) service.AuthService {
	wire.Build(
		mysql.NewUserRepo,
		implement.NewAuthService,
	)
	return nil
}

func initUcCreateUser(db *gorm.DB) user.CreateUser {
	wire.Build(
		initUserService,
		user.NewCreateUser,
	)

	return nil
}

func initUcUpdateUser(db *gorm.DB) user.UpdateUser {
	wire.Build(
		initUserService,
		user.NewUpdateUser,
	)

	return nil
}

func initUcDeleteUser(db *gorm.DB) user.DeleteUser {
	wire.Build(
		initUserService,
		user.NewDeleteUser,
	)

	return nil
}

func initUcGetUser(db *gorm.DB) user.GetUser {
	wire.Build(
		initUserService,
		user.NewGetUser,
	)

	return nil
}

func initUcLogin(db *gorm.DB, cfg *config.Config) auth.Login {
	wire.Build(
		InitAuthService,
		initUserService,
		auth.NewLogin,
	)

	return nil
}

func initUcAuthInit(db *gorm.DB, cfg *config.Config) auth.AuthInit {
	wire.Build(
		InitAuthService,
		auth.NewAuthInit,
	)

	return nil
}
func InitLoginHandler(db *gorm.DB, cfg *config.Config) func(ctx *gin.Context) {
	wire.Build(
		initUcLogin,
		doLogin,
	)
	return nil
}
func InitCreateUserHandler(db *gorm.DB, cfg *config.Config) func(ctx *gin.Context) {
	wire.Build(
		initUcCreateUser,
		createUser,
	)
	return nil
}

func InitUpdateUserHandler(db *gorm.DB, cfg *config.Config) func(ctx *gin.Context) {
	wire.Build(
		initUcUpdateUser,
		updateUser,
	)
	return nil
}

func InitDeleteUserHandler(db *gorm.DB, cfg *config.Config) func(ctx *gin.Context) {
	wire.Build(
		initUcDeleteUser,
		deleteUser,
	)
	return nil
}

func InitGetUserHandler(db *gorm.DB, cfg *config.Config) func(ctx *gin.Context) {
	wire.Build(
		initUcGetUser,
		getUser,
	)
	return nil
}

func InitAuthInitHandler(db *gorm.DB, cfg *config.Config) func(ctx *gin.Context) {
	wire.Build(
		initUcAuthInit,
		doAuthInit,
	)
	return nil
}

func initUcListUser(db *gorm.DB) user.ListUser {
	wire.Build(
		initUserService,
		user.NewListUser,
	)

	return nil
}
