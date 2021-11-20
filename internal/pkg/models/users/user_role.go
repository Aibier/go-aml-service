package users

import (
	"github.com/Aibier/go-aml-service/internal/pkg/models"
	"gorm.io/gorm"
	"time"
)

// UserRole ...
type UserRole struct {
	models.Model
	UserID   uint64 `gorm:"column:user_id;unique_index:user_role;not null;" json:"user_id"`
	RoleName string `gorm:"column:role_name;not null;" json:"role_name"`
}

// BeforeCreate ...
func (m *UserRole) BeforeCreate(*gorm.DB) error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate ...
func (m *UserRole) BeforeUpdate(*gorm.DB) error {
	m.UpdatedAt = time.Now()
	return nil
}
