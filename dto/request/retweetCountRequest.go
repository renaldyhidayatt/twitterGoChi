package request

import "github.com/go-playground/validator/v10"

type RetweetCountRequest struct {
	UserID  int `json:"user_id" validate:"required"`
	TweetID int `json:"tweet_id" validate:"required"`
	TweetBy int `json:"tweetby" validate:"required"`
}

func (r *RetweetCountRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
