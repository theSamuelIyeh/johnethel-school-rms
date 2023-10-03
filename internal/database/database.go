package database

import (
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

// database interface
type Database interface {
	CreateUser() error
	CheckForUser(username string, role string) (*User, error)
	StoreRefreshToken(token string, userId string, expTime time.Time) error
	ClearRefreshToken(userId string) error
	GetUserById(id string, role string) (*User, error)
	ResetPassword(email_or_admission_no string, role string, password string, passwordhash string) (*User, error)
}

// database struct
type DatabaseImpl struct {
	Engine *xorm.Engine
}

// database engine instance
var Engine *xorm.Engine

// initialize database connection
func InitDB() error {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	var err error
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", username, password, host, port, name)
	Engine, err = xorm.NewEngine("postgres", connStr)
	if err != nil {
		return err
	}
	log.Print("Database connection established")
	err = Engine.Sync()
	if err != nil {
		return err
	}
	return nil
}

// close database connection
func CloseDB() {
	if Engine != nil {
		Engine.Close()
	}
}

// create database struct instance
func NewDatabase(engine *xorm.Engine) *DatabaseImpl {
	return &DatabaseImpl{Engine: engine}
}
