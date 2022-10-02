package implement

import (
	"go_web/domain/entity"
	"go_web/domain/repository"
	"go_web/domain/service"
	"go_web/errors"

	"gorm.io/gorm"
)

type userServiceImpl struct {
	userRepo repository.UserRepo
}

var errDuplicateEmail = errors.BadRequestFromStr("duplicate email")

func NewUserService(userRep repository.UserRepo) service.UserService {
	return &userServiceImpl{
		userRepo: userRep,
	}
}

func (u *userServiceImpl) WithTx(tx *gorm.DB) service.UserService {
	return &userServiceImpl{
		userRepo: u.userRepo.WithTx(tx),
	}
}
func (u *userServiceImpl) GetByEmail(email string) (result *entity.User, err error) {
	return u.userRepo.FindByEmail(email)
}

func (u *userServiceImpl) GetByID(ID uint32) (result *entity.User, err error) {
	return u.userRepo.FindByID(ID)
}

func (u *userServiceImpl) DeleteByID(ID uint32) error {
	return u.userRepo.DeleteByID(ID)
}
func (u *userServiceImpl) Create(user entity.User) (result *entity.User, err error) {
	return u.userRepo.Create(user)
}

func (u *userServiceImpl) Update(ID uint32, user entity.User) (result *entity.User, err error) {
	return u.userRepo.Update(ID, user)
}

func (u *userServiceImpl) ValidateCreate(user entity.User) error {
	if err := user.Validate(); err != nil {
		return err
	}
	count, err := u.userRepo.CountByEmail(user.Email)
	if err != nil {
		return err
	}
	if count > 0 {
		return errDuplicateEmail
	}
	return nil
}

func (u *userServiceImpl) ValidateUpdate(user entity.User) error {
	if err := user.Validate(); err != nil {
		return err
	}
	count, err := u.userRepo.CountByEmailExcludeID(user.Email, user.ID)
	if err != nil {
		return err
	}

	if count > 0 {
		return errDuplicateEmail
	}

	return nil
}
