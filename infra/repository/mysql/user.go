package mysql

import (
	"go_web/domain/entity"
	"go_web/domain/repository"

	"gorm.io/gorm"
)

type userRepoImpl struct {
	db *gorm.DB
	repository.BaseRepo[entity.User]
}

func NewUserRepo(db *gorm.DB) repository.UserRepo {
	return &userRepoImpl{
		db,
		newBaseRepoImpl(db),
	}
}

func (e *userRepoImpl) WithTx(tx *gorm.DB) repository.UserRepo {
	return NewUserRepo(tx)
}

func (e *userRepoImpl) CountByEmail(email string) (int64, error) {
	return e.CountByConds("email = ?", email)
}

func (e *userRepoImpl) CountByEmailExcludeID(email string, ID uint32) (int64, error) {
	return e.CountByConds("email = ? AND ID != ?", email, ID)
}
