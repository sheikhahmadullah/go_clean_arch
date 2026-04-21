package usecase

import "golang-rest-api/domain"

type TaskUsecase struct {
	repo domain.TaskRepository
}

func NewTaskUsecase(r domain.TaskRepository) *TaskUsecase {
	return &TaskUsecase{r}
}

func (u *TaskUsecase) CreateTask(title string, userID uint) error {
	task := domain.Task{
		Title:  title,
		UserId: userID,
	}

	return u.repo.Create(&task)
}

func (u *TaskUsecase) GetTasks(userID uint) ([]domain.Task, error) {
	return u.repo.GetByUserID(userID)
}
