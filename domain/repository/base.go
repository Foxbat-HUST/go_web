package repository

import (
	"go_web/domain/entity"
	"go_web/infra/model"
)

type OrderDirectionType string

func (o OrderDirectionType) IsValid() bool {
	return o == ASC || o == DESC
}

const (
	ASC  OrderDirectionType = "ASC"
	DESC OrderDirectionType = "DESC"
)

type Order struct {
	ColumnName string
	Direction  OrderDirectionType
}
type Condition struct {
	Clause string
	Value  any
}
type GetListOptions struct {
	PageIndex   *int
	ItemPerPage *int
	OrderBy     []Order
	Conditions  []Condition
}
type BaseRepo[M model.Model, E entity.Entity] interface {
	GetList(options GetListOptions) ([]*E, int64, error)
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
