package storage 

import "TideUp/internal/models"

type ContextStorage interface {
	CreateContext(newContext *models.Context) error
	DeleteContext(contextID int) error 
	EditContext(contextID int,newContext models.Context) error 
	ShowAllContexts(limit int) ([]models.Context, error) 
}

func (s *Storage) CreateContext(newContext *models.Context) error {
	err := s.db.Create(&newContext).Error
	return err
}

func (s *Storage) DeleteContext(contextID int) error {
	err := s.db.Delete(&contextID).Error
	return err
}

func (s *Storage) ShowAllContexts(limit int) ([]models.Context,error) {
	var allContexts []models.Context
	err := s.db.Limit(limit).Find(&allContexts).Error
	return allContexts,err
}

func (s *Storage) EditContext(contextID int,newContext models.Context) error {
	return s.db.Model(models.Context{}).
	Where("id = ?", contextID).
	Updates(newContext).Error
}

