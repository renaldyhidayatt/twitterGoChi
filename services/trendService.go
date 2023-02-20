package services

import (
	db "github.com/renaldyhidayatt/twittersqlc/db/sqlc"
	"github.com/renaldyhidayatt/twittersqlc/dto/request"
	"github.com/renaldyhidayatt/twittersqlc/interfaces"
	"github.com/renaldyhidayatt/twittersqlc/repository"
)

type TrendService = interfaces.ITrendService

type trendService struct {
	repository repository.TrendRepository
}

func NewTrendService(repository repository.TrendRepository) *trendService {
	return &trendService{repository: repository}
}

func (s *trendService) GetTrend() ([]db.GetTrendsRow, error) {
	res, err := s.repository.GetTrend()

	return res, err
}

func (s *trendService) GetTrendByHash(hashtag string) ([]string, error) {
	res, err := s.repository.GetTrendByHash(hashtag)

	return res, err

}

func (s *trendService) CreateTrend(req request.TrendRequest) (db.Trend, error) {
	res, err := s.repository.CreateTrend(req)

	return res, err
}
