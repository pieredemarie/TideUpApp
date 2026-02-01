package task

import (
	"TideUp/internal/dto"
	"TideUp/internal/models"
	"TideUp/internal/storage"
	"time"
)

type TaskService struct {
	Storage storage.TaskStorage
}

type ITaskService interface {
	CreateTask(task models.Task) error
	RemoveTask(userID int, taskID int) error
	ShowAllTasks(userID, limit int) ([]models.Task, error)
	UpdateTask(userID int, taskID int, req dto.UpdateTaskRequest) error

	//GetTasksFromEbbMode returns tasks scheduled for current day only
	GetTasksFromEbbMode(userID int) ([]models.Task, error)

	// GetFloatingTasks returns tasks without deadline
	GetFloatingTasks(userID int) ([]models.Task, error)
}

func NewTaskService(storage storage.TaskStorage) *TaskService {
	return &TaskService{
		Storage: storage,
	}
}

func (s *TaskService) RemoveTask(userID int, taskID int) error {
	return s.Storage.RemoveTask(userID, taskID)
}

func (s *TaskService) CreateTask(task models.Task) error {
	return s.Storage.AddTask(&task)
}

func (s *TaskService) ShowAllTasks(userID, limit int) ([]models.Task, error) {
	return s.Storage.ShowAllTasks(userID, limit)
}

func (s *TaskService) UpdateTask(userID int, taskID int, req dto.UpdateTaskRequest) error {
	return s.Storage.UpdateTask(userID, taskID, req)
}

func (s *TaskService) GetTasksFromEbbMode(userID int) ([]models.Task, error) {
	today := time.Now().Truncate(24 * time.Hour)
	return s.Storage.GetTasksByDate(userID, today)
}

func (s *TaskService) GetFloatingTasks(userID int) ([]models.Task, error) {
	return s.Storage.GetTasksWithDeadlineNull(userID)
}
