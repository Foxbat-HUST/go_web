package model

import (
	"go_web/infra/model/internal/gorm/gen"
)

type Job struct {
	gen.Job
	BaseModel
}
