package repository

import (
	"context"
	"fmt"
	"time"

	db "github.com/renaldyhidayatt/twittersqlc/db/sqlc"
	"github.com/renaldyhidayatt/twittersqlc/dto/request"
	"github.com/renaldyhidayatt/twittersqlc/interfaces"
)

type CommentRepository = interfaces.ICommentRepository

type commentRepository struct {
	db  *db.Queries
	ctx context.Context
}

func NewCommentRepository(db *db.Queries, ctx context.Context) *commentRepository {
	return &commentRepository{db: db, ctx: ctx}
}

func (r *commentRepository) Comment(req request.CreateCommentRequest) (db.Comment, error) {
	var commentDb db.Comment

	res, err := r.WasCommentBy(req.CommentBy, req.CommentOn)

	if err != nil {
		return db.Comment{}, fmt.Errorf("failed was comment : %w", err)
	}

	if res == commentDb {
		if req.CommentBy != req.PostedBy {
			err = r.db.DeleteNotification(r.ctx, db.DeleteNotificationParams{
				NotificationFor:  int32(req.PostedBy),
				NotificationFrom: int32(req.CommentBy),
				Target:           int32(req.CommentOn),
				Type:             "comment",
			})

			if err != nil {
				return db.Comment{}, fmt.Errorf("failed error :%w", err)
			}
		}
	} else {
		if req.CommentBy != req.PostedBy {
			_, err := r.db.CreateNotification(r.ctx, db.CreateNotificationParams{
				NotificationFor:   int32(req.PostedBy),
				NotificationFrom:  int32(req.CommentBy),
				Target:            int32(req.CommentOn),
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
			CommentBy: int32(req.CommentBy),
			CommentOn: int32(req.CommentOn),
			Comment:   req.Comment,
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

func (r *commentRepository) DeleteComment(req request.DeleteCommentRequest) (bool, error) {
	var comment db.Comment

	res, err := r.WasCommentBy(req.CommentBy, req.CommentOn)

	if err != nil {
		return false, fmt.Errorf("failed :%w", err)
	}

	if res == comment {
		if req.CommentBy != req.TweetBy {
			err = r.db.DeleteNotification(r.ctx, db.DeleteNotificationParams{
				NotificationFor:  int32(req.TweetBy),
				NotificationFrom: int32(req.CommentBy),
				Target:           int32(req.CommentOn),
				Type:             "comment",
			})

			if err != nil {
				return false, fmt.Errorf("failed :%w", err)
			}
		}

		err = r.db.DeleteComment(r.ctx, int32(req.CommentID))

		if err != nil {
			return true, fmt.Errorf("failed :%w", err)
		}
	}

	return true, nil
}

func (r *commentRepository) RepliesTweet(email string) ([]db.RepliesTweetsRow, error) {
	res, err := r.db.GetUsernameOREmail(r.ctx, email)

	if err != nil {
		return nil, fmt.Errorf("failed user undefined: %w", err)
	}

	resemail, err := r.db.RepliesTweets(r.ctx, int32(res.UserID))

	if err != nil {
		return nil, fmt.Errorf("failed error :%w", err)
	}

	return resemail, nil
}
