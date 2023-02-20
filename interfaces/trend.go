package interfaces

import (
	db "github.com/renaldyhidayatt/twittersqlc/db/sqlc"
	"github.com/renaldyhidayatt/twittersqlc/dto/request"
)

type ITrendRepository interface {
	GetTrend() ([]db.GetTrendsRow, error)
	GetTrendByHash(hashtag string) ([]string, error)
	CreateTrend(req request.TrendRequest) (db.Trend, error)
}

type ITrendService interface {
	GetTrend() ([]db.GetTrendsRow, error)
	GetTrendByHash(hashtag string) ([]string, error)
	CreateTrend(req request.TrendRequest) (db.Trend, error)
}
