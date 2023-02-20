package request

import "github.com/go-playground/validator/v10"

type CreateCommentRequest struct {
	CommentBy int    `json:"commentBy" validate:"required"`
	CommentOn int    `json:"commentOn" validate:"required"`
	Comment   string `json:"comment" validate:"required"`
	PostedBy  int    `json:"postedBy" validate:"required"`
}

func (c *CreateCommentRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
