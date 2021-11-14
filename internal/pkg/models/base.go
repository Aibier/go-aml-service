package models

import (
	"github.com/google/uuid"
	"time"
)

type Model struct {
	UUID      uuid.UUID `gorm:"column:uuid;not null;" json:"uuid" form:"uuid"`
	ID        uint64    `gorm:"column:id;primary_key;auto_increment;" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;not null;" json:"updated_at"`
}
