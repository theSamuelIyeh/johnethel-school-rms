package services

import (
	"github.com/thesamueliyeh/johnethel-rms/internal/database"
	"golang.org/x/crypto/bcrypt"
)

// credentials service struct
type credentialsService struct {
	db database.Database
}

// credentials service factory
func NewCredentialsService(db database.Database) *credentialsService {
	return &credentialsService{db: db}
}

// credentials service verify credentials method
func (s *credentialsService) VerifyCredentials(admission_no string, password string, role string) (*database.User, error) {
	user, err := s.db.CheckForUser(admission_no, role)
	if err != nil {
		return nil, err
	}
	passwordHash := user.Password
	if passwordHash == "" {
		passwordHash = user.ResetPassword
	}
	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		return nil, err
	}
	return user, nil
}
