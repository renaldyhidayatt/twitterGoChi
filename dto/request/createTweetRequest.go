package request

import "github.com/go-playground/validator/v10"

type CreateTweetRequest struct {
	Status     string `json:"status"`
	TweetBy    int32  `json:"tweetBy"`
	TweetImage string `json:"tweetImage"`
}

func (r *CreateTweetRequest) Validate() error {
	validate := validator.New()

	return validate.Struct(r)
}
