package request

import "github.com/go-playground/validator/v10"

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=5,max=20"`
	Password string `json:"password" validate:"required,min=6"`
}

func (req *LoginRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(req)
}
