package storage

import "TideUp/internal/models"

type TaskStorage interface {
	AddTask(newTask *models.Task) error
	RemoveTask(taskID int) error 
	UpdateTask(taskID int, newTask *models.Task) error
	ShowAllTasks(limit int) ([]models.Task, error)
	MakeTaskFloat(taskID int) error
}


func (s *Storage) AddTask(newTask *models.Task) error {
	err := s.db.Create(newTask).Error
	return err
}

func (s *Storage) RemoveTask(taskID int) error {
	err := s.db.
	Where("id = ?",taskID).
	Delete(&models.Task{}).Error
	return err
}

func (s *Storage) ShowAllTasks(limit int) ([]models.Task, error) {
	var tasks []models.Task
	err := s.db.Limit(limit).Find(&tasks).Error
	return tasks,err
}

func (s *Storage) UpdateTask(taskID int, newTask *models.Task) error {
	return s.db.Model(&models.Task{}).
	Where("id = ?", taskID).
	Updates(newTask).Error
}

func (s *Storage) MakeTaskFloat(taskID int) error {
	return nil
}
