package services

import (
	db "github.com/renaldyhidayatt/twittersqlc/db/sqlc"
	"github.com/renaldyhidayatt/twittersqlc/dto/request"
	"github.com/renaldyhidayatt/twittersqlc/interfaces"
	"github.com/renaldyhidayatt/twittersqlc/repository"
)

type CommentService = interfaces.ICommentService

type commentService struct {
	repository repository.CommentRepository
}

func NewCommentService(repository repository.CommentRepository) *commentService {
	return &commentService{repository: repository}
}

func (s *commentService) RepliesTweet(email string) ([]db.RepliesTweetsRow, error) {
	res, err := s.repository.RepliesTweet(email)

	return res, err
}

func (s *commentService) Comment(req request.CreateCommentRequest) (db.Comment, error) {
	res, err := s.repository.Comment(req)

	return res, err
}

func (s *commentService) DeleteComment(req request.DeleteCommentRequest) (bool, error) {
	res, err := s.repository.DeleteComment(req)

	return res, err
}
