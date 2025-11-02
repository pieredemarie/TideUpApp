package storage

import (
	"TideUp/internal/dto"
	"TideUp/internal/models"
)

type ContextStorage interface {
	CreateContext(newContext *models.Context) error
	DeleteContext(userID,contextID int) error 
	EditContext(userID,contextID int,newContext dto.UpdateContextRequest) error 
	ShowAllContexts(userID, limit int) ([]models.Context, error) 
}

func (s *Storage) CreateContext(newContext *models.Context) error {
	err := s.db.Create(&newContext).Error
	return err
}

func (s *Storage) DeleteContext(userID,contextID int) error {
	return s.db.Where("id = ? AND user_id = ?", contextID, userID).Delete(&models.Context{}).Error
}

func (s *Storage) ShowAllContexts(userID,limit int) ([]models.Context,error) {
	var allContexts []models.Context
	err := s.db.Where("user_id = ?",userID).Limit(limit).Find(&allContexts).Error
	return allContexts,err
}

func (s *Storage) EditContext(userID,contextID int,newContext dto.UpdateContextRequest) error {
	return s.db.Model(models.Context{}).
	Where("id = ? AND user_id = ?", contextID, userID).
	Updates(newContext).Error
}

