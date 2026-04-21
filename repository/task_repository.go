package repository

import (
	"golang-rest-api/domain"

	"gorm.io/gorm"
)

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) domain.TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) Create(task *domain.Task) error {
	return r.db.Create(task).Error
}

func (r *taskRepository) GetByUserID(userID uint) ([]domain.Task, error) {
	var tasks []domain.Task

	err := r.db.Where("user_id = ?", userID).Find(&tasks).Error

	return tasks, err
}
