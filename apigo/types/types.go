package types

import "time"

type UserStore interface{
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
}

type User struct{
	ID 				int 			`json:"id"`
	FirstName string 		`json:"firstname"`
	LastName 	string 		`json:"lastname"`
	Email 		string 		`json:"email"`
	Password  string    `json:"-"` //não queremos que a senha seja serializada
	CreatedAt time.Time `json:"created_at"`
}

type RegisterUser struct{
	Email 		string 	`json:"email" validate:"required,email"`	
	Password 	string 	`json:"password" validate:"required,min=6,max=100"`
	FirstName string 	`json:"firstname" validate:"required"`
	LastName 	string 	`json:"lastname" validate:"required"`
}

