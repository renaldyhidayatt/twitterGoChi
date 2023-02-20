package request

import "github.com/go-playground/validator/v10"

type LikesRequest struct {
	LikedBy int `json:"likedby" validate:"required"`
	TweetID int `json:"tweet_id" validate:"required"`
	TweetBy int `json:"tweetby" validate:"required"`
}

func (r *LikesRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
