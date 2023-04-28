package mysql

import (
	"go_web/domain/entity"
	"go_web/domain/repository"
	"go_web/errors"
	"go_web/infra/model"
	"strings"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

var invalidParam = errors.InternalServerErrFromStr("invalid query pram")
var invalidGetListOption = errors.BadRequestFromStr("invalid get list options")

type baseRepo[M model.Model, E entity.Entity] struct {
	db *gorm.DB
}

func newBaseRepoImpl[M model.Model, E entity.Entity](db *gorm.DB) repository.BaseRepo[M, E] {
	return &baseRepo[M, E]{
		db: db,
	}
}

func (e *baseRepo[M, E]) GetList(options repository.GetListOptions) ([]*E, int64, error) {
	if !e.isValidGetListOption(options) {
		return nil, 0, invalidGetListOption
	}
	var model M
	query := e.db.Model(&model)
	for _, cond := range options.Conditions {
		query.Where(cond.Clause, cond.Value)
	}
	var count int64
	err := query.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	if options.PageIndex != nil {
		offset := *options.ItemPerPage * (*options.PageIndex - 1)
		limit := *options.ItemPerPage
		query.Offset(offset).Limit(limit)
	}

	rawResults := []*M{}
	if err = query.Find(&rawResults).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, errors.NotFound(err)
		}

		return nil, 0, err
	}

	results := []*E{}
	if err = copy(&results, &rawResults); err != nil {
		return nil, 0, err
	}

	return results, count, nil
}

func (e *baseRepo[M, E]) Create(params E) (*E, error) {
	var model M
	if err := copy(&model, &params); err != nil {
		return nil, err
	}

	if err := e.db.Create(model).Error; err != nil {
		return nil, err
	}

	var result E
	if err := copier.Copy(&result, &model); err != nil {
		return nil, err
	}

	return &result, nil
}

func (e *baseRepo[M, E]) Update(ID uint32, params E) (*E, error) {
	var model M
	if err := e.db.First(&model, ID).Error; err != nil {
		return nil, err
	}

	if err := copyIgnoreEmpty(&model, &params); err != nil {
		return nil, err
	}

	if err := e.db.Save(model).Error; err != nil {
		return nil, err
	}

	var result E
	if err := copy(&result, &model); err != nil {
		return nil, err
	}

	return &result, nil
}

func (e *baseRepo[M, E]) FindByID(ID uint32) (*E, error) {
	var model M
	if err := e.db.First(&model, ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NotFound(err)
		}
		return nil, err
	}

	var result E
	if err := copy(&result, &model); err != nil {
		return nil, err
	}

	return &result, nil
}

func (e *baseRepo[M, E]) FindByIDs(IDs []uint32) ([]E, error) {
	models := []M{}
	if err := e.db.Find(&models, IDs).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NotFound(err)
		}
		return nil, err
	}

	result := []E{}
	if err := copy(&result, &models); err != nil {
		return nil, err
	}

	return result, nil
}

func (e *baseRepo[M, E]) FindOneByConds(conds string, params ...interface{}) (*E, error) {
	if strings.Count(conds, "?") != len(params) {
		return nil, invalidParam
	}
	var model M
	if err := e.db.First(&model, conds, params).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NotFound(err)
		}
		return nil, err
	}

	var result E
	if err := copy(&result, &model); err != nil {
		return nil, err
	}

	return &result, nil
}

func (e *baseRepo[M, E]) FindAllByConds(conds string, params ...interface{}) ([]E, error) {
	if strings.Count(conds, "?") != len(params) {
		return nil, invalidParam
	}

	models := []M{}
	if err := e.db.Find(&models, conds, params).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NotFound(err)
		}
		return nil, err
	}

	result := []E{}
	if err := copy(&result, &models); err != nil {
		return nil, err
	}

	return result, nil
}

func (e *baseRepo[M, E]) CountByConds(conds string, params ...interface{}) (int64, error) {
	if strings.Count(conds, "?") != len(params) {
		return 0, invalidParam
	}
	var count int64
	var model M
	if err := e.db.Model(&model).Where(conds, params...).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (e *baseRepo[M, E]) DeleteByID(ID uint32) error {
	var model M
	return e.db.Where("ID = ?", ID).Delete(&model).Error
}

func (e *baseRepo[M, E]) DeleteByIDs(IDs []uint32) (int64, error) {
	var model M
	result := e.db.Where("ID IN ?", IDs).Delete(&model)
	return result.RowsAffected, result.Error
}

func (e *baseRepo[M, E]) isValidGetListOption(option repository.GetListOptions) bool {
	if option.PageIndex != nil && option.ItemPerPage == nil {
		return false
	}

	// pageIndex start from 1...
	if option.PageIndex != nil && *option.PageIndex < 1 {
		return false
	}

	var model M
	for _, sort := range option.OrderBy {
		if !sort.Direction.IsValid() {
			return false
		}

		if _, exist := model.Columns()[sort.ColumnName]; !exist {
			return false
		}

	}
	return true
}

func copy(toValue interface{}, fromValue interface{}) (err error) {
	return copier.Copy(toValue, fromValue)
}

func copyIgnoreEmpty(toValue interface{}, fromValue interface{}) (err error) {
	return copier.CopyWithOption(toValue, fromValue, copier.Option{
		IgnoreEmpty: true,
	})
}
