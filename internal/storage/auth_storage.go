package storage

import (
	"TideUp/internal/models"
)

type IAuth interface {
	CreateUser(newUser *models.User) error 
	GetUserByEmail(email string) (*models.User,error)
	GetUserPassword(email string) (string, error)
}

func (s *Storage) CreateUser(newUser *models.User) error  {
	return s.db.Create(newUser).Error
}

func (s *Storage) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := s.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (s *Storage) GetUserPassword(email string) (string, error) {
	var user models.User
	err := s.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return "", err
	}

	return user.PasswordHash, nil
}