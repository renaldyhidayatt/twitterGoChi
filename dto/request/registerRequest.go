package request

import "github.com/go-playground/validator/v10"

type RegisterRequest struct {
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
	Username  string `json:"username" validate:"required,min=5,max=20"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
}

func (req *RegisterRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(req)
}
