package request

import "github.com/go-playground/validator/v10"

type CreateRetweetRequest struct {
	RetweetBy   int32 `json:"retweetBy" validate:"required"`
	RetweetFrom int32 `json:"retweetFrom" validate:"required"`
}

func (r *CreateRetweetRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
