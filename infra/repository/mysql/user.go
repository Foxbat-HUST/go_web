package mysql

import (
	"go_web/domain/entity"
	"go_web/domain/repository"
	"go_web/errors"
	"go_web/infra/model"

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
func (e *userRepoImpl) FindAuthUserByEmail(email string) (*entity.AuthUser, error) {
	var rawResult model.User
	if err := e.db.First(&rawResult, "email = ?", email).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NotFound(err)
		}
		return nil, err
	}

	result := &entity.AuthUser{}
	if err := copy(result, &rawResult); err != nil {
		return nil, err
	}

	return result, nil
}

func (e *userRepoImpl) FindByEmail(email string) (*entity.User, error) {
	var rawResult model.User
	if err := e.db.First(&rawResult, "email = ?", email).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NotFound(err)
		}
		return nil, err
	}

	result := &entity.User{}
	if err := copy(result, &rawResult); err != nil {
		return nil, err
	}

	return result, nil
}
