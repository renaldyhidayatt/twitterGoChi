package request

import "github.com/go-playground/validator/v10"

type DeleteCommentRequest struct {
	CommentID int `json:"commentid" validate:"required"`
	CommentBy int `json:"commentby" validate:"required"`
	CommentOn int `json:"commenton" validate:"required"`
	TweetBy   int `json:"tweetby" validate:"required"`
}

func (d *DeleteCommentRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}
