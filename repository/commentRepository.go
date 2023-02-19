package repository

import (
	"context"
	"fmt"
	"time"

	db "github.com/renaldyhidayatt/twittersqlc/db/sqlc"
)

type commentRepository struct {
	db  *db.Queries
	ctx context.Context
}

func NewCommentRepository(db *db.Queries, ctx context.Context) *commentRepository {
	return &commentRepository{db: db, ctx: ctx}
}

func (r *commentRepository) Comment(commentBy int, commentOn int, comment string, postedBy int) (db.Comment, error) {
	var commentDb db.Comment

	res, err := r.WasCommentBy(commentBy, commentOn)

	if err != nil {
		return db.Comment{}, fmt.Errorf("failed was comment : %w", err)
	}

	if res == commentDb {
		if commentBy != postedBy {
			err = r.db.DeleteNotification(r.ctx, db.DeleteNotificationParams{
				NotificationFor:  int32(postedBy),
				NotificationFrom: int32(commentBy),
				Target:           int32(commentOn),
				Type:             "comment",
			})

			if err != nil {
				return db.Comment{}, fmt.Errorf("failed error :%w", err)
			}
		}
	} else {
		if commentBy != postedBy {
			_, err := r.db.CreateNotification(r.ctx, db.CreateNotificationParams{
				NotificationFor:   int32(postedBy),
				NotificationFrom:  int32(commentBy),
				Target:            int32(commentOn),
				Type:              "comment",
				Status:            0,
				NotificationCount: 0,
				NotificationOn:    time.Now(),
			})

			if err != nil {
				return db.Comment{}, fmt.Errorf("failed create notification :%w", err)
			}
		}
		res, err := r.db.CreateComment(r.ctx, db.CreateCommentParams{
			CommentBy: int32(commentBy),
			CommentOn: int32(commentOn),
			Comment:   comment,
			CommentAt: time.Now(),
		})

		if err != nil {
			return db.Comment{}, fmt.Errorf("failed error :%w", err)
		}

		return res, nil

	}
	return db.Comment{}, nil
}

func (r *commentRepository) WasCommentBy(commentBy int, commentOn int) (db.Comment, error) {
	res, err := r.db.WasCommentBy(r.ctx, db.WasCommentByParams{
		CommentOn: int32(commentOn),
		CommentBy: int32(commentBy),
	})

	if err != nil {
		return db.Comment{}, fmt.Errorf("failed was comment : %w", err)
	}

	return res, nil
}

func (r *commentRepository) DeleteComment(comment_id int, commentBy int, commentOn int, tweetby int) (bool, error) {
	var comment db.Comment

	res, err := r.WasCommentBy(commentBy, commentOn)

	if err != nil {
		return false, fmt.Errorf("failed :%w", err)
	}

	if res == comment {
		if commentBy != tweetby {
			err = r.db.DeleteNotification(r.ctx, db.DeleteNotificationParams{
				NotificationFor:  int32(tweetby),
				NotificationFrom: int32(commentBy),
				Target:           int32(commentOn),
				Type:             "comment",
			})

			if err != nil {
				return false, fmt.Errorf("failed :%w", err)
			}
		}

		err = r.db.DeleteComment(r.ctx, int32(comment_id))

		if err != nil {
			return true, fmt.Errorf("failed :%w", err)
		}
	}

	return true, nil
}
