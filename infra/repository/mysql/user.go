package mysql

import (
	"go_web/domain/entity"
	"go_web/domain/repository"
	"go_web/infra/repository/mysql/model"

	"gorm.io/gorm"
)

type userRepoImpl struct {
	db   *gorm.DB
	base baseRepo[model.User, entity.User]
}

func NewUserRepo(db *gorm.DB) repository.UserRepo {
	return &userRepoImpl{
		db: db,
		base: baseRepo[model.User, entity.User]{
			db: db,
		},
	}
}

func (e *userRepoImpl) WithTx(tx *gorm.DB) repository.UserRepo {
	return NewUserRepo(tx)
}

func (e *userRepoImpl) Create(param entity.User) (*entity.User, error) {
	return e.base.create(param)
}

func (e *userRepoImpl) Update(ID uint32, param entity.User) (*entity.User, error) {
	return e.base.update(ID, param)
}

func (e *userRepoImpl) FindByID(ID uint32) (*entity.User, error) {
	return e.base.findByID(ID)
}

func (e *userRepoImpl) FindByIDs(IDs []uint32) ([]entity.User, error) {
	return e.base.findByIDs(IDs)
}

func (e *userRepoImpl) DeleteByID(ID uint32) error {
	return e.base.deleteByID(ID)
}

func (e *userRepoImpl) DeleteByIDs(IDs []uint32) (int64, error) {
	return e.base.deleteByIDs(IDs)
}

func (e *userRepoImpl) CountByEmail(email string) (int64, error) {
	return e.base.countByConds("email = ?", email)
}

func (e *userRepoImpl) CountByEmailExcludeID(email string, ID uint32) (int64, error) {
	return e.base.countByConds("email = ? AND ID != ?", email, ID)
}
