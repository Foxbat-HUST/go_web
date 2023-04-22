package model

import (
	"go_web/infra/model/internal/gorm/gen"

	"gorm.io/gorm"
)

type User struct {
	gen.User
	gorm.Model
}
