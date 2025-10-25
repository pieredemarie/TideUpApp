package auth

import (
	"TideUp/internal/models"
	"TideUp/internal/storage"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Storage storage.IAuth
}

func NewAuthService(storage storage.IAuth) *AuthService {
	return &AuthService{
		Storage: storage,
	}
}

func (s *AuthService) Register(email,name,password string) error {
	_, err := s.Storage.GetUserByEmail(email)
	if err == nil {
		return errors.New("email already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.User{
		Name: name,
		Email: email,
		PasswordHash: string(hash),
	}

	return s.Storage.CreateUser(*user)
} 

func (s *AuthService) Login(email,password string) (string, error)  {
	user, err := s.Storage.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("invalid password or email")
	} 	

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash),[]byte(password))
	if err != nil {
		return "", errors.New("invalid password or email")
	}

	token, err := GenerateToken(user.ID)
	if err != nil {
		return "", errors.New("couldn't generate token")
	}
	
	return token, nil
}