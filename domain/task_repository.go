package domain

type TaskRepository interface {
	Create(task *Task) error
	GetByUserID(userID uint) ([]Task, error)
}
