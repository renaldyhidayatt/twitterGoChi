package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	db "github.com/renaldyhidayatt/twittersqlc/db/sqlc"
)

type retweetRepository struct {
	db  *db.Queries
	ctx context.Context
}

func NewRetweetRepository(db *db.Queries, ctx context.Context) *retweetRepository {
	return &retweetRepository{db: db, ctx: ctx}
}

func (r *retweetRepository) RetweetCount(user_id int, tweet_id int, tweet_by int) (string, error) {
	res, err := r.WasRetweetBy(user_id, tweet_id)

	if err != nil {
		return "", fmt.Errorf("failed: %w", err)
	}

	if res {
		if user_id != tweet_id {
			err = r.db.DeleteNotification(r.ctx, db.DeleteNotificationParams{
				NotificationFor:  int32(tweet_by),
				NotificationFrom: int32(user_id),
				Target:           int32(tweet_id),
				Type:             db.EnumNotificationTypeRetweet,
			})
			if err != nil {
				return "", fmt.Errorf("failed to delete notification: %w", err)
			}
		}

		_, err = r.DeleteRetweet(user_id, tweet_id)

		if err != nil {
			return "", fmt.Errorf("failed to delete retweet: %w", err)
		}

		result := map[string]interface{}{
			"retweet": -1,
		}
		response, err := json.Marshal(result)
		if err != nil {
			return "", fmt.Errorf("failed to marshal response: %w", err)
		}

		return string(response), nil
	} else {
		if user_id != tweet_by {
			_, err = r.db.CreateNotification(r.ctx, db.CreateNotificationParams{
				NotificationFor:   int32(tweet_by),
				NotificationFrom:  int32(user_id),
				Target:            int32(tweet_id),
				Type:              "retweet",
				Status:            0,
				NotificationCount: 0,
				NotificationOn:    time.Now(),
			})
			if err != nil {
				return "", fmt.Errorf("failed to create notification: %w", err)
			}
		}

		_, err = r.db.CreateRetweet(r.ctx, db.CreateRetweetParams{
			RetweetBy:   int32(user_id),
			RetweetFrom: int32(tweet_id),
		})
		if err != nil {
			return "", fmt.Errorf("failed to create retweet: %w", err)
		}

		result := map[string]interface{}{
			"retweet": 1,
		}
		response, err := json.Marshal(result)
		if err != nil {
			return "", fmt.Errorf("failed to marshal response: %w", err)
		}

		return string(response), nil
	}

}

func (r *retweetRepository) WasRetweetBy(user_id int, tweet_id int) (bool, error) {
	var retweet db.WasRetweetByParams

	retweet.RetweetBy = int32(user_id)    // user_id
	retweet.RetweetFrom = int32(tweet_id) // tweet_id

	_, err := r.db.WasRetweetBy(r.ctx, retweet)

	if err != nil {
		return false, fmt.Errorf("failed wasretweetby : %w", err)
	}

	return true, nil

}

func (r *retweetRepository) DeleteRetweet(user_id int, tweet_id int) (bool, error) {
	err := r.db.DeleteRetweet(r.ctx, db.DeleteRetweetParams{
		RetweetBy:   int32(user_id),
		RetweetFrom: int32(tweet_id),
	})

	if err != nil {
		return false, fmt.Errorf("failed delete retweet : %w", err)
	}

	return true, nil
}

func (r *retweetRepository) ResetRetweetCount(user_id int, tweet_id int, tweetby int) (string, error) {
	res, err := r.WasRetweetBy(user_id, tweet_id)

	if err != nil {
		return "", fmt.Errorf("failed to check retweet: %w", err)
	}

	if res {
		if user_id != tweetby {
			err = r.db.DeleteNotification(r.ctx, db.DeleteNotificationParams{
				NotificationFor:  int32(tweetby),
				NotificationFrom: int32(user_id),
				Target:           int32(tweet_id),
				Type:             "retweet",
			})

			if err != nil {
				return "", fmt.Errorf("failed to delete retweet: %w", err)
			}

			result := map[string]int{"deretweet": -1}
			jsonResult, err := json.Marshal(result)

			if err != nil {
				return "", fmt.Errorf("failed to marshal json result: %w", err)
			}

			return string(jsonResult), nil
		}
	}

	return "", nil
}

func (r *retweetRepository) GetRetweet(tweet_id int) (int64, error) {
	res, err := r.db.GetRetweet(r.ctx, int32(tweet_id))

	if err != nil {
		return 0, fmt.Errorf("")
	}

	return res, nil
}
