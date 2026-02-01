package context

import (
	"TideUp/internal/dto"
	"TideUp/internal/models"
	"TideUp/internal/storage"
)

type contextService struct {
	Storage storage.ContextStorage
}

func NewContextService(storage storage.ContextStorage) *contextService {
	return &contextService{
		Storage: storage,
	}
}

type ContextService interface {
	Create(newContext *models.Context) error
	Delete(userID, contextID int) error
	Edit(userID, contextID int, newContext dto.UpdateContextRequest) error
	ShowAll(userID, limit int) ([]models.Context, error)
}

func (s *contextService) Create(newContext *models.Context) error {
	return s.Storage.CreateContext(newContext)
}

func (s *contextService) Delete(userID, contextID int) error {
	return s.Storage.DeleteContext(userID, contextID)
}

func (s *contextService) Edit(userID, contextID int, newContext dto.UpdateContextRequest) error {
	return s.Storage.EditContext(userID, contextID, newContext)
}

func (s *contextService) ShowAll(userID, limit int) ([]models.Context, error) {
	return s.Storage.ShowAllContexts(userID, limit)
}
