package request

import "github.com/go-playground/validator/v10"

type UnFollowRequest struct {
	UnfollowID int `json:"unfollow_id"`
	UserID     int `json:"user_id"`
}

func (req *UnFollowRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(req)
}
