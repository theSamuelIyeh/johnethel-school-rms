package services

import "github.com/thesamueliyeh/johnethel-rms/internal/database"

// services
var CredentialsService *credentialsService
var TokenService *tokenService
var UserService *userService
var RedirectService *redirectService
var EmailService *emailService

func InitServices(db database.Database) {
	CredentialsService = NewCredentialsService(db)
	TokenService = NewTokenService(db)
	UserService = NewUserService(db)
	RedirectService = NewRedirectService(db)
	EmailService = NewEmailService(db)
}
