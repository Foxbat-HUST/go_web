package model

import (
	"go_web/infra/model/internal/gorm/gen"
)

type User struct {
	gen.User
	BaseModel
}
