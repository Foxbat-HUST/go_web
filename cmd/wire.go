//go:build exclude

package cmd

import (
	"go_web/api/usecase/user"
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
