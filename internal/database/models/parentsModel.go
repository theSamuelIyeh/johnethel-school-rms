package models

type Parent struct {
	Id      string `xorm:"'id' <-"`
	User_id string `xorm:"'user_id'"`
	Email   string `xorm:"'email'"`
}

func (*Parent) TableName() string {
	return "parents"
}

// ---------------- SQL VERSION ----------------

// CREATE TABLE "parents" (
// 	"id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
// 	"user_id" uuid,
// 	"email" varchar UNIQUE NOT NULL
//   );
