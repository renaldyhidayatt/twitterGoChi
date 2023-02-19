package repository

import (
	"context"
	"fmt"
	"time"

	db "github.com/renaldyhidayatt/twittersqlc/db/sqlc"
	"github.com/renaldyhidayatt/twittersqlc/dto/request"
)

type trendRepository struct {
	db  *db.Queries
	ctx context.Context
}

func NewTrendRepository(db *db.Queries, ctx context.Context) *trendRepository {
	return &trendRepository{db: db, ctx: ctx}
}

func (r *trendRepository) GetTrend() ([]db.GetTrendsRow, error) {
	res, err := r.db.GetTrends(r.ctx)

	if err != nil {
		return nil, fmt.Errorf("failed error trends: %w", err)
	}

	return res, nil
}

func (r *trendRepository) GetTrendByHash(hashtag string) ([]string, error) {
	res, err := r.db.GetTrendByHash(r.ctx, hashtag)

	if err != nil {
		return nil, fmt.Errorf("failed error trend hashtag: %w", err)
	}

	return res, nil
}

func (r *trendRepository) CreateTrend(req request.TrendRequest) (db.Trend, error) {
	var trendCreate db.CreateTrendParams

	trendCreate.Hashtag = req.Hashtag
	trendCreate.TweetId = req.TweetId
	trendCreate.UserID = req.UserID
	trendCreate.CreatedOn = time.Now()

	res, err := r.db.CreateTrend(r.ctx, trendCreate)

	if err != nil {
		return db.Trend{}, fmt.Errorf("failed error create trend: %w", err)
	}

	return res, nil
}
