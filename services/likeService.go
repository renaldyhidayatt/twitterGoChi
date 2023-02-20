package services

import (
	"github.com/renaldyhidayatt/twittersqlc/dto/request"
	"github.com/renaldyhidayatt/twittersqlc/interfaces"
	"github.com/renaldyhidayatt/twittersqlc/repository"
)

type LikesService = interfaces.ILikesService

type likeService struct {
	repository repository.LikeRepository
}

func NewLikeService(repository repository.LikeRepository) *likeService {
	return &likeService{repository: repository}
}

func (s *likeService) GetLikes(tweet_id int) (int64, error) {
	res, err := s.repository.GetLikes(tweet_id)

	return res, err
}

func (s *likeService) Likes(req request.LikesRequest) (map[string]int, error) {
	res, err := s.repository.Likes(req)

	return res, err
}
