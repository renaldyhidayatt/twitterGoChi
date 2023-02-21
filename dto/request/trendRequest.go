package request

import "github.com/go-playground/validator/v10"

type TrendRequest struct {
	Hashtag string `json:"hashtag"`
	TweetId int32  `json:"tweetId"`
	UserID  int32  `json:"user_id"`
}

func (r *TrendRequest) Validate() error {
	validate := validator.New()

	return validate.Struct(r)
}
