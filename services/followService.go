package services

import (
	db "github.com/renaldyhidayatt/twittersqlc/db/sqlc"
	"github.com/renaldyhidayatt/twittersqlc/dto/request"
	"github.com/renaldyhidayatt/twittersqlc/interfaces"
	"github.com/renaldyhidayatt/twittersqlc/repository"
)

type FollowService = interfaces.IFollowService

type followService struct {
	repository repository.FollowRepository
}

func NewFollowService(repository repository.FollowRepository) *followService {
	return &followService{repository: repository}
}

func (s *followService) CheckFollow(req request.FollowCheckRequest) (db.Follow, error) {
	res, err := s.repository.CheckFollow(req)

	return res, err
}

func (s *followService) WhoToFollow(email string) ([]db.User, error) {
	res, err := s.repository.WhoToFollow(email)

	return res, err
}

func (s *followService) Follow(req request.AddFollowRequest) (db.ResultFollowOrUnFollowRow, error) {
	res, err := s.repository.Follow(req)

	return res, err
}

func (s *followService) UnFollow(req request.UnFollowRequest) (db.ResultFollowOrUnFollowRow, error) {
	res, err := s.repository.UnFollow(req)

	return res, err
}

func (s *followService) ResultFollowingList(email string) ([]db.ResultFollowingListRow, error) {
	res, err := s.repository.ResultFollowingList(email)

	return res, err
}

func (s *followService) ResultFollowersList(email string) ([]db.ResultFollowersListRow, error) {
	res, err := s.repository.ResultFollowersList(email)

	return res, err
}

func (s *followService) SuggestedList(email string) ([]db.SuggestedListRow, error) {
	res, err := s.repository.SuggestedList(email)

	return res, err
}
