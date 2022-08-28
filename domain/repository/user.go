package repository

import (
	"go_web/domain/entity"

	"gorm.io/gorm"
)

type UserRepo interface {
	WithTx(tx *gorm.DB) UserRepo
	Create(param entity.User) (*entity.User, error)
	Update(ID uint32, param entity.User) (*entity.User, error)
	FindByID(ID uint32) (*entity.User, error)
	FindByIDs(IDs []uint32) ([]entity.User, error)
	DeleteByID(ID uint32) error
	DeleteByIDs(IDs []uint32) (int64, error)
	CountByEmail(email string) (int64, error)
	CountByEmailExcludeID(email string, ID uint32) (int64, error)
}
