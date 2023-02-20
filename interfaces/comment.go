package interfaces

import (
	db "github.com/renaldyhidayatt/twittersqlc/db/sqlc"
	"github.com/renaldyhidayatt/twittersqlc/dto/request"
)

type ICommentRepository interface {
	RepliesTweet(email string) ([]db.RepliesTweetsRow, error)
	Comment(req request.CreateCommentRequest) (db.Comment, error)
	DeleteComment(req request.DeleteCommentRequest) (bool, error)
}

type ICommentService interface {
	RepliesTweet(email string) ([]db.RepliesTweetsRow, error)
	Comment(req request.CreateCommentRequest) (db.Comment, error)
	DeleteComment(req request.DeleteCommentRequest) (bool, error)
}
