package request

import "github.com/go-playground/validator/v10"

type GetTweetRequest struct {
	TweetID int `json:"tweet_id" validate:"required"`
	TweetBy int `json:"tweet_by" validate:"required"`
}

func (r *GetTweetRequest) Validate() error {
	validate := validator.New()

	return validate.Struct(r)
}
