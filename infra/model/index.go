package model

import (
	"time"

	"gorm.io/gorm"
)

type Model interface {
	Columns() map[string]bool
	User | Job
}

type BaseModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
