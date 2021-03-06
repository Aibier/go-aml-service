package persistence

import (
	"strconv"

	"github.com/Aibier/go-aml-service/internal/pkg/db"
	models "github.com/Aibier/go-aml-service/internal/pkg/models/users"
)

// UserRepository ...
type UserRepository struct{}

var userRepository *UserRepository

// GetUserRepository ...
func GetUserRepository() *UserRepository {
	if userRepository == nil {
		userRepository = &UserRepository{}
	}
	return userRepository
}

// Get ...
func (r *UserRepository) Get(id string) (*models.User, error) {
	var user models.User
	where := models.User{}
	where.ID, _ = strconv.ParseUint(id, 10, 64)
	_, err := First(&where, &user, []string{"Role"})
	if err != nil {
		return nil, err
	}
	return &user, err
}

// GetByUsername ...
func (r *UserRepository) GetByUsername(username string) (*models.User, error) {
	var user models.User
	where := models.User{}
	where.Username = username
	_, err := First(&where, &user, []string{"Role"})
	if err != nil {
		return nil, err
	}
	return &user, err
}

// All ...
func (r *UserRepository) All() (*[]models.User, error) {
	var users []models.User
	err := Find(&models.User{}, &users, []string{"Role"}, "id asc")
	return &users, err
}

// Query ...
func (r *UserRepository) Query(q *models.User) (*[]models.User, error) {
	var users []models.User
	err := Find(&q, &users, []string{"Role"}, "id asc")
	return &users, err
}

// Add ...
func (r *UserRepository) Add(user *models.User) error {
	err := Create(&user)
	if err != nil {
		return err
	}
	err = Save(&user)
	if err != nil {
		return err
	}
	return err
}

// Update ...
func (r *UserRepository) Update(user *models.User) error {
	var userRole models.UserRole
	_, err := First(models.UserRole{UserID: user.ID}, &userRole, []string{})
	if err != nil {
		return err
	}
	userRole.RoleName = user.Role.RoleName
	err = Save(&userRole)
	if err != nil {
		return err
	}
	err = db.GetDB().Omit("Role").Save(&user).Error
	if err != nil {
		return err
	}
	user.Role = userRole
	return err
}

// Delete ...
func (r *UserRepository) Delete(user *models.User) error {
	err := db.GetDB().Unscoped().Delete(models.UserRole{UserID: user.ID}).Error
	if err != nil {
		return err
	}
	err = db.GetDB().Unscoped().Delete(&user).Error
	if err != nil {
		return err
	}
	return err
}
