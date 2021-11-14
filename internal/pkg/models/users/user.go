package users

import (
	"github.com/Aibier/go-aml-service/internal/pkg/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	models.Model
	Username  string   `gorm:"column:username;not null;unique_index:username" json:"username" form:"username"`
	Firstname string   `gorm:"column:firstname;not null;" json:"firstname" form:"firstname"`
	Lastname  string   `gorm:"column:lastname;not null;" json:"lastname" form:"lastname"`
	Hash      string   `gorm:"column:hash;not null;" json:"hash"`
	Role      UserRole `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (m *User) BeforeCreate(*gorm.DB) error {
	m.UUID = uuid.New()
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

func (m *User) BeforeUpdate(*gorm.DB) error {
	m.UpdatedAt = time.Now()
	return nil
}
