package repository

import (
	"context"
	"fmt"
	"time"

	db "github.com/renaldyhidayatt/twittersqlc/db/sqlc"
)

type likeRepository struct {
	db  *db.Queries
	ctx context.Context
}

func NewLikeRepository(db *db.Queries, ctx context.Context) *likeRepository {
	return &likeRepository{db: db, ctx: ctx}
}

func (r *likeRepository) GetLikes(tweet_id int) (int64, error) {
	res, err := r.db.GetLikes(r.ctx, int32(tweet_id))

	if err != nil {
		return 0, fmt.Errorf("error failed getlikes: %w", err)
	}

	return res, nil
}

func (r *likeRepository) Likes(likedBy int, tweetId int, tweetBy int) (map[string]int, error) {
	likes, err := r.db.WasLikeBy(r.ctx, db.WasLikeByParams{
		LikeBy: int32(likedBy),
		LikeOn: int32(tweetId),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to check if tweet was liked by user: %w", err)
	}

	result := make(map[string]int)
	if likes.LikeID == 0 {
		// Not liked
		if likedBy != tweetBy {
			if _, err := r.db.CreateNotification(r.ctx, db.CreateNotificationParams{
				NotificationFor:   int32(tweetBy),
				NotificationFrom:  int32(likedBy),
				Type:              "like",
				Status:            0,
				NotificationCount: 0,
				NotificationOn:    time.Now(),
			}); err != nil {
				return nil, fmt.Errorf("failed to create notification: %w", err)
			}
		}
		if _, err := r.db.CreateLike(r.ctx, db.CreateLikeParams{
			LikeBy: int32(likedBy),
			LikeOn: int32(tweetId),
		}); err != nil {
			return nil, fmt.Errorf("failed to create like: %w", err)
		}
		result["likes"] = 1
	} else {
		// Liked
		if likedBy != tweetBy {
			if err := r.db.DeleteNotification(r.ctx, db.DeleteNotificationParams{
				NotificationFor:  int32(tweetBy),
				NotificationFrom: int32(likedBy),
				Target:           int32(tweetId),
				Type:             "like",
			}); err != nil {
				return nil, fmt.Errorf("failed to delete notification: %w", err)
			}
		}
		if err := r.db.DeleteLike(r.ctx, db.DeleteLikeParams{
			LikeBy: int32(likedBy),
			LikeOn: int32(tweetId),
		}); err != nil {
			return nil, fmt.Errorf("failed to delete like: %w", err)
		}
		result["likes"] = -1
	}

	return result, nil
}
