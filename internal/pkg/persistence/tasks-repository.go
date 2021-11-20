package persistence

import (
	"github.com/Aibier/go-aml-service/internal/pkg/db"
	models "github.com/Aibier/go-aml-service/internal/pkg/models/tasks"
	"strconv"
)
// TaskRepository ...
type TaskRepository struct{}

var taskRepository *TaskRepository

// GetTaskRepository ...
func GetTaskRepository() *TaskRepository {
	if taskRepository == nil {
		taskRepository = &TaskRepository{}
	}
	return taskRepository
}

// Get ...
func (r *TaskRepository) Get(id string) (*models.Task, error) {
	var task models.Task
	where := models.Task{}
	where.ID, _ = strconv.ParseUint(id, 10, 64)
	_, err := First(&where, &task, []string{"User"})
	if err != nil {
		return nil, err
	}
	return &task, err
}

// All ...
func (r *TaskRepository) All() (*[]models.Task, error) {
	var tasks []models.Task
	err := Find(&models.Task{}, &tasks, []string{"User"}, "id asc")
	return &tasks, err
}

// Query ...
func (r *TaskRepository) Query(q *models.Task) (*[]models.Task, error) {
	var tasks []models.Task
	err := Find(&q, &tasks, []string{"User"}, "id asc")
	return &tasks, err
}

// Add ...
func (r *TaskRepository) Add(task *models.Task) error {
	err := Create(&task)
	if err != nil {
		return err
	}
	err = Save(&task)
	if err != nil {
		return err
	}
	return err
}

// Update ...
func (r *TaskRepository) Update(task *models.Task) error {
	return db.GetDB().Omit("User").Save(&task).Error
}

// Delete ...
func (r *TaskRepository) Delete(task *models.Task) error {
	return db.GetDB().Unscoped().Delete(&task).Error
}
