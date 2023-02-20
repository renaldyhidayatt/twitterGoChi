package request

import "github.com/go-playground/validator/v10"

type ResetRetweetRequest struct {
	UserID  int `json:"user_id" validate:"required"`
	TweetID int `json:"tweet_id" validate:"required"`
	TweetBy int `json:"tweet_by" validate:"required"`
}

func (r *ResetRetweetRequest) Validate() error {

	validate := validator.New()
	return validate.Struct(r)
}
