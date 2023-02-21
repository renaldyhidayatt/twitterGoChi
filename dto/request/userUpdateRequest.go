package request

import "github.com/go-playground/validator/v10"

type UpdateUserRequest struct {
	Username     string `json:"username"`
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	Username_2   string `json:"username_2"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Profileimage string `json:"profileimage"`
	Profilecover string `json:"profilecover"`
	Bio          string `json:"bio"`
	Country      string `json:"country"`
	Website      string `json:"website"`
}

func (r *UpdateUserRequest) Validate() error {
	validate := validator.New()

	return validate.Struct(r)
}
