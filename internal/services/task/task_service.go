package task

import (
	"TideUp/internal/models"
	"TideUp/internal/storage"
	"time"
)

type TaskService struct {
	Storage storage.TaskStorage
}

func NewTaskService(storage storage.TaskStorage) *TaskService {
	return &TaskService{
		Storage: storage,
	}
}

func (s *TaskService) GetTasksFromEbbMode(userID int) ([]models.Task, error) {
	today := time.Now().Truncate(24* time.Hour)
	return s.Storage.GetTasksByDate(userID,today)
}

func (s *TaskService) GetFloatingTasks(userID int) ([]models.Task, error) {
	return s.Storage.GetTasksWithDeadlineNull(userID)
}

