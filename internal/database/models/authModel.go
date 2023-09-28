package models

import (
	"time"
)

type Auth struct {
	Id           string    `xorm:"'id' <-"`
	UserId       string    `xorm:"'user_id'"`
	RefreshToken string    `xorm:"'refreshtoken'"`
	ExpiryDate   time.Time `xorm:"'expirydate'"`
}

func (*Auth) TableName() string {
	return "auth"
}

// ---------------- SQL VERSION ----------------

// CREATE TABLE "auth" (
//   "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
//   "user_id" uuid,
//   "refreshtoken" varchar UNIQUE NOT NULL,
//   "expirydate" timestamp NOT NULL
// );
