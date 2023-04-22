package model

import (
	"time"

	"gorm.io/gorm"
)

type Model interface {
	User
}

type BaseModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
