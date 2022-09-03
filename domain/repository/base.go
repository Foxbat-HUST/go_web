package repository

import (
	"go_web/domain/entity"
)

type BaseRepo[E entity.Entity] interface {
	Create(param E) (*E, error)
	Update(ID uint32, param E) (*E, error)
	FindByID(ID uint32) (*E, error)
	FindByIDs(IDs []uint32) ([]E, error)
	FindOneByConds(conds string, params ...interface{}) (*E, error)
	FindAllByConds(conds string, params ...interface{}) ([]E, error)
	CountByConds(conds string, params ...interface{}) (int64, error)
	DeleteByID(ID uint32) error
	DeleteByIDs(IDs []uint32) (int64, error)
}
