package repository

import (
	"context"
	"fmt"

	db "github.com/renaldyhidayatt/twittersqlc/db/sqlc"
	"github.com/renaldyhidayatt/twittersqlc/dto/request"
	"github.com/renaldyhidayatt/twittersqlc/interfaces"
)

type TweetRepository = interfaces.ITweetRepository

type tweetRepository struct {
	db  *db.Queries
	ctx context.Context
}

func NewTweetRepository(db *db.Queries, ctx context.Context) *tweetRepository {
	return &tweetRepository{db: db, ctx: ctx}
}

func (r *tweetRepository) Tweets(email string) ([]db.GetTweetAllRow, error) {
	res_user, err := r.db.GetUsernameOREmail(r.ctx, email)

	if err != nil {
		return nil, fmt.Errorf("failed result user: %w", err)
	}

	res, err := r.db.GetTweetAll(r.ctx, int32(res_user.UserID))

	if err != nil {
		return nil, fmt.Errorf("failed results tweet: %w", err)
	}

	return res, nil
}

func (r *tweetRepository) TweetByMe(email string) ([]db.GetTweetByMeRow, error) {
	res_user, err := r.db.GetUsernameOREmail(r.ctx, email)

	if err != nil {
		return nil, fmt.Errorf("failed result user: %w", err)
	}

	res, err := r.db.GetTweetByMe(r.ctx, int32(res_user.UserID))

	if err != nil {
		return nil, fmt.Errorf("failed results tweet: %w", err)
	}

	return res, nil
}

func (r *tweetRepository) GetHashTagTweet(hashtag string) ([]db.GetHashTagTweetRow, error) {
	res, err := r.db.GetHashTagTweet(r.ctx, hashtag)

	if err != nil {
		return nil, fmt.Errorf("failed result hashtag: %w", err)
	}

	return res, nil
}

func (r *tweetRepository) GetMention(email string) ([]db.User, error) {
	resuser, err := r.db.GetUsernameOREmail(r.ctx, email)

	if err != nil {
		return nil, fmt.Errorf("failed result user: %w", err)
	}

	res, err := r.db.GetMention(r.ctx, resuser.Email)

	if err != nil {
		return nil, fmt.Errorf("failed result mention: %w", err)
	}

	return res, nil
}

func (r *tweetRepository) TweetCounts(tweet_by int) (int64, error) {
	res, err := r.db.TweetCounts(r.ctx, int32(tweet_by))

	if err != nil {
		return 0, fmt.Errorf("failed result tweetcount :%w", err)
	}

	return res, nil
}

func (r *tweetRepository) GetTweet(req request.GetTweetRequest) (db.GetTweetRow, error) {
	res, err := r.db.GetTweet(r.ctx, db.GetTweetParams{
		TweetID: int32(req.TweetID),
		TweetBy: int32(req.TweetBy),
	})

	if err != nil {
		return db.GetTweetRow{}, fmt.Errorf("failed error :%w", err)
	}

	return res, nil
}

func (r *tweetRepository) CreateTweet(req request.CreateTweetRequest) (db.Tweet, error) {
	res, err := r.db.CreateTweet(r.ctx, db.CreateTweetParams{
		Status:     req.Status,
		TweetBy:    req.TweetBy,
		TweetImage: req.TweetImage,
	})

	if err != nil {
		return db.Tweet{}, fmt.Errorf("failed error :%w", err)
	}

	return res, nil
}
