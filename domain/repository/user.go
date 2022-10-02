package repository

import (
	"go_web/domain/entity"

	"gorm.io/gorm"
)

type UserRepo interface {
	BaseRepo[entity.User]
	WithTx(tx *gorm.DB) UserRepo
	CountByEmail(email string) (int64, error)
	CountByEmailExcludeID(email string, ID uint32) (int64, error)
	FindAuthUserByEmail(email string) (*entity.AuthUser, error)
	FindByEmail(email string) (*entity.User, error)
}
