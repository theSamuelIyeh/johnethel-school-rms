package database

import (
	"fmt"
	"time"

	"github.com/thesamueliyeh/johnethel-rms/internal/database/models"
)

// create user method
func (db *DatabaseImpl) CreateUser() error {
	var err error
	return err
}

type User struct {
	models.User   `xorm:"extends"`
	models.Pupil  `xorm:"extends"`
	models.Staff  `xorm:"extends"`
	models.Parent `xorm:"extends"`
}

// Check for user method
func (db *DatabaseImpl) CheckForUser(admission_no_or_email string, role string) (*User, error) {
	var user User
	var found bool
	var err error
	if role == "pupil" {
		found, err = db.Engine.Table("users").Alias("u").Join("INNER", "pupils p", "p.user_id = u.id").Where("p.admission_no = ?", admission_no_or_email).Get(&user)
	} else if role == "staff" {
		found, err = db.Engine.Table("users").Alias("u").Join("INNER", "staffs s", "s.user_id = u.id").Where("s.admission_no = ?", admission_no_or_email).Get(&user)
	} else if role == "parent" {
		found, err = db.Engine.Table("users").Alias("u").Join("INNER", "parents p", "p.user_id = u.id").Where("p.email = ?", admission_no_or_email).Get(&user)
	} else {
		return nil, fmt.Errorf("invalid role")
	}
	if err != nil {
		return nil, err
	} else if !found {
		return nil, fmt.Errorf("user not found")
	}
	return &user, nil
}

// store refresh token in db method
func (db *DatabaseImpl) StoreRefreshToken(token string, userId string, expTime time.Time) error {
	found, err := db.Engine.Table("auth").Where("user_id = ?", userId).Exist()
	if err != nil {
		return err
	} else if found {
		affected, err := db.Engine.Table("auth").Where("user_id = ?", userId).Update(models.Auth{RefreshToken: token, ExpiryDate: expTime})
		if err != nil {
			return err
		} else if affected == 0 {
			return fmt.Errorf("couldn't update token")
		}
		return nil
	}
	affected, err := db.Engine.Table("auth").Insert(models.Auth{RefreshToken: token, UserId: userId, ExpiryDate: expTime})
	if err != nil {
		return err
	} else if affected == 0 {
		return fmt.Errorf("couldn't insert token")
	}
	return nil
}

// clear refresh token in db method
func (db *DatabaseImpl) ClearRefreshToken(userId string) error {
	affected, err := db.Engine.Table("auth").Where("user_id = ?", userId).Delete(models.Auth{})
	if err != nil {
		return err
	} else if affected == 0 {
		return fmt.Errorf("couldn't delete token")
	}
	return nil
}

// get user from db method
func (db *DatabaseImpl) GetUserById(id string, role string) (*User, error) {
	var user User
	var found bool
	var err error
	if role == "pupil" {
		found, err = db.Engine.Table("users").Alias("u").Where("u.id = ?", id).Join("INNER", "pupils p", "p.user_id = u.id").Get(&user)
	} else if role == "staff" {
		found, err = db.Engine.Table("users").Alias("u").Where("u.id = ?", id).Join("INNER", "staffs s", "s.user_id = u.id").Get(&user)
	} else if role == "parent" {
		found, err = db.Engine.Table("users").Alias("u").Where("u.id = ?", id).Join("INNER", "parents p", "p.user_id = u.id").Get(&user)
	} else {
		return nil, fmt.Errorf("invalid role")
	}
	if err != nil {
		return nil, err
	} else if !found {
		return nil, fmt.Errorf("user not found")
	}
	return &user, nil
}
