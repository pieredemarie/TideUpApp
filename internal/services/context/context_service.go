package context

import "TideUp/internal/storage"

type ContextService struct {
	Storage storage.ContextStorage
}

func NewContextService(storage storage.ContextStorage) *ContextService {
	return &ContextService{
		Storage: storage,
	}
}