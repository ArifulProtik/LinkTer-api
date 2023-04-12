package utility

import "github.com/google/uuid"

type UserInput struct {
	Name        string `validate:"required"`
	Username    string `validate:"required,min=6"`
	Email       string `validate:"required,email"`
	Password    string `validate:"required,min=8"`
	RefPassword string `validate:"required,min=8"`
}

type SessionInput struct {
	UserID uuid.UUID
	Token  string
	IP     string
}
type SignInInput struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}
