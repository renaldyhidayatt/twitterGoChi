package services

import (
	"github.com/renaldyhidayatt/twittersqlc/dto/request"
	"github.com/renaldyhidayatt/twittersqlc/interfaces"
	"github.com/renaldyhidayatt/twittersqlc/repository"
)

type RetweetService = interfaces.IRetweetRepository

type retweetService struct {
	repository repository.RetweetRepository
}

func NewRetweetService(repository repository.RetweetRepository) *retweetService {
	return &retweetService{repository: repository}
}

func (s *retweetService) RetweetCount(req request.RetweetCountRequest) (string, error) {
	res, err := s.repository.RetweetCount(req)

	return res, err
}

func (s *retweetService) ResetRetweetCount(req request.ResetRetweetRequest) (string, error) {
	res, err := s.repository.ResetRetweetCount(req)

	return res, err
}

func (s *retweetService) GetRetweet(tweet_id int) (int64, error) {
	res, err := s.repository.GetRetweet(tweet_id)

	return res, err
}
