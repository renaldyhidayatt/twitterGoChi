package repository

import (
	"context"
	"fmt"

	db "github.com/renaldyhidayatt/twittersqlc/db/sqlc"
)

type tweetRepository struct {
	db  *db.Queries
	ctx context.Context
}

func NewTweetRepository(db *db.Queries, ctx context.Context) *tweetRepository {
	return &tweetRepository{db: db, ctx: ctx}
}

func (r *tweetRepository) Tweets(user_id int) ([]db.GetTweetAllRow, error) {
	res, err := r.db.GetTweetAll(r.ctx, int32(user_id))

	if err != nil {
		return nil, fmt.Errorf("failed results tweet: %w", err)
	}

	return res, nil
}

func (r *tweetRepository) TweetByMe(user_id int) ([]db.GetTweetByMeRow, error) {
	res, err := r.db.GetTweetByMe(r.ctx, int32(user_id))

	if err != nil {
		return nil, fmt.Errorf("failed results tweet: %w", err)
	}

	return res, nil
}
