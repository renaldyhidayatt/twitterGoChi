package request

import "github.com/go-playground/validator/v10"

type AddFollowRequest struct {
	Sender   int32 `json:"sender"`
	Receiver int32 `json:"receiver"`
}

func (d *AddFollowRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}
