package model

import (
	"time"
)

type BaseModel struct {
	ID        int       `gorm:"column:id;primary_key" json:"id"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"-"`
	CreatedAt time.Time `gorm:"column:created_at" json:"-"`
}

func (base BaseModel) IsPersisted() bool {
	return base.ID > 0
}
