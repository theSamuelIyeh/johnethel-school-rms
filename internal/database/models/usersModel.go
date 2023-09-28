package models

type User struct {
	Id                string `xorm:"'id' <-"`
	Password          string `xorm:"'password'"`
	ResetPassword     string `xorm:"'resetpassword'"`
	ResetPasswordRaw  string `xorm:"'resetpasswordraw'"`
	PasswordSalt      string `xorm:"'passwordsalt'"`
	ResetPasswordDone bool   `xorm:"'resetpassworddone'"`
	ProfilePic        string `xorm:"'profilepic'"`
	Role              string `xorm:"'role'"`
	FirstName         string `xorm:"'firstname'"`
	Middlename        string `xorm:"'middlename'"`
	LastName          string `xorm:"'lastname'"`
	CreatedAt         string `xorm:"'created_at'"`
	UpdatedAt         string `xorm:"'updated_at'"`
	DOB               string `xorm:"'dob'"`
	Gender            string `xorm:"gender"`
	LastLogin         string `xorm:"'last_login'"`
	HomeAddress       string `xorm:"'homeaddress'"`
	Active            bool   `xorm:"'active'"`
}

func (*User) TableName() string {
	return "users"
}

// ---------------- SQL VERSION ----------------

// CREATE TABLE "users" (
// 	"id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
// 	"password" varchar NOT NULL DEFAULT '',
// 	"resetpassword" varchar NOT NULL DEFAULT '',
// 	"resetpasswordraw" varchar NOT NULL DEFAULT '',
// 	"passwordsalt" varchar NOT NULL DEFAULT '',
// 	"resetpassworddone" boolean NOT NULL DEFAULT false,
// 	"profilepic" varchar NOT NULL DEFAULT 'default.jpg',
// 	"role" roles NOT NULL,
// 	"firstname" varchar NOT NULL,
// 	"middlename" varchar NOT NULL,
// 	"lastname" varchar NOT NULL,
// 	"created_at" timestamp NOT NULL,
// 	"updated_at" timestamp NOT NULL,
// 	"dob" timestamp NOT NULL,
// 	"gender" sex NOT NULL,
// 	"last_login" timestamp NOT NULL,
// 	"homeaddress" text NOT NULL,
// 	"active" boolean NOT NULL DEFAULT false
//   );
