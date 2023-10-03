package services

import (
	"github.com/thesamueliyeh/johnethel-rms/internal/database"
	"golang.org/x/crypto/bcrypt"
)

// user service struct
type userService struct {
	db database.Database
}

// user service factory
func NewUserService(db database.Database) *userService {
	return &userService{db: db}
}

// user service get user by id method
func (s *userService) GetUserById(id string, role string) (*database.User, error) {
	user, err := s.db.GetUserById(id, role)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) ResetPassword(email string, role string) (*database.User, error) {
	resetpassword := ""
	resetpasswordhash, err := bcrypt.GenerateFromPassword([]byte(resetpassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user, err := s.db.ResetPassword(email, role, resetpassword, string(resetpasswordhash))
	if err != nil {
		return nil, err
	}
	return user, nil
}
