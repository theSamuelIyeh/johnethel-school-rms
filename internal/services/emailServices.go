package services

import "github.com/thesamueliyeh/johnethel-rms/internal/database"

type emailService struct {
	db database.Database
}

func NewEmailService(db database.Database) *emailService {
	return &emailService{db: db}
}

func (s *emailService) SendResetPasswordEmail(user *database.User) (bool, error) {
	return true, nil
}
