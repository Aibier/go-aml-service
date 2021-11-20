package tasks

import (
	"github.com/Aibier/go-aml-service/internal/pkg/models"
	"github.com/Aibier/go-aml-service/internal/pkg/models/users"
	"gorm.io/gorm"
	"time"
)

// Task ...
type Task struct {
	models.Model
	Name   string     `gorm:"column:name;not null;" json:"name" form:"name"`
	Text   string     `gorm:"column:text;not null;" json:"text" form:"text"`
	UserID uint64     `gorm:"column:user_id;unique_index:user_id;not null;" json:"user_id" form:"user_id"`
	User   users.User `json:"user"`
}

// BeforeCreate ...
func (m *Task) BeforeCreate(*gorm.DB) error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate ...
func (m *Task) BeforeUpdate(*gorm.DB) error {
	m.UpdatedAt = time.Now()
	return nil
}
