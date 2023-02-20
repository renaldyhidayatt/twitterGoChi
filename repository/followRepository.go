package repository

import (
	"context"
	"fmt"
	"time"

	db "github.com/renaldyhidayatt/twittersqlc/db/sqlc"
	"github.com/renaldyhidayatt/twittersqlc/dto/request"
	"github.com/renaldyhidayatt/twittersqlc/interfaces"
)

type FollowRepository = interfaces.IFollowRepository

type followRepository struct {
	db  *db.Queries
	ctx context.Context
}

func NewFollowRepository(db *db.Queries, ctx context.Context) *followRepository {
	return &followRepository{db: db, ctx: ctx}
}

func (r *followRepository) CheckFollow(req request.FollowCheckRequest) (db.Follow, error) {
	var followCheck db.CheckFollowParams

	followCheck.Sender = int32(req.UserID)
	followCheck.Receiver = int32(req.Following)

	res, err := r.db.CheckFollow(r.ctx, followCheck)

	if err != nil {
		return db.Follow{}, fmt.Errorf("failed error check follow")
	}

	return res, nil
}

func (r *followRepository) WhoToFollow(email string) ([]db.User, error) {
	res_usr, err := r.db.GetUsernameOREmail(r.ctx, email)

	if err != nil {
		return nil, fmt.Errorf("failed undefined user: %w", err)
	}

	res, err := r.db.WhoToFollow(r.ctx, int32(res_usr.UserID))

	if err != nil {
		return nil, fmt.Errorf("failed whotofollow: %w", err)
	}

	return res, nil
}

func (r *followRepository) Follow(req request.AddFollowRequest) (db.ResultFollowOrUnFollowRow, error) {
	var addFollow db.AddFollowParams

	addFollow.Sender = req.Sender
	addFollow.Receiver = req.Receiver
	addFollow.FollowStatus = "1"
	addFollow.FollowOn = time.Now()

	res, err := r.db.AddFollow(r.ctx, addFollow)

	if err != nil {
		return db.ResultFollowOrUnFollowRow{}, err
	}

	var result db.ResultFollowOrUnFollowParams

	result.Sender = res.Sender
	result.UserID = res.Sender

	_, err = r.db.AddFollowingCount(r.ctx, req.Sender)

	if err != nil {
		return db.ResultFollowOrUnFollowRow{}, nil
	}

	_, err = r.db.AddFollowerCount(r.ctx, req.Sender)

	if err != nil {
		return db.ResultFollowOrUnFollowRow{}, nil
	}

	ress, err := r.db.ResultFollowOrUnFollow(r.ctx, result)

	if err != nil {
		return db.ResultFollowOrUnFollowRow{}, fmt.Errorf("failed follow :%w", err)
	}

	return ress, nil
}

func (r *followRepository) UnFollow(req request.UnFollowRequest) (db.ResultFollowOrUnFollowRow, error) {
	var unfollow db.UnFollowParams

	unfollow.Sender = int32(req.UserID)
	unfollow.Receiver = int32(req.UnfollowID)

	res, err := r.db.UnFollow(r.ctx, unfollow)

	if err != nil {
		return db.ResultFollowOrUnFollowRow{}, fmt.Errorf("faild unfollow :%w", err)
	}

	_, err = r.db.RemoveFollowingCount(r.ctx, int32(req.UserID))

	if err != nil {
		return db.ResultFollowOrUnFollowRow{}, fmt.Errorf("faild removefollowing :%w", err)
	}

	_, err = r.db.RemoveFollowersCount(r.ctx, int32(req.UserID))

	if err != nil {
		return db.ResultFollowOrUnFollowRow{}, fmt.Errorf("failed removefollowers: %w", err)
	}

	var result db.ResultFollowOrUnFollowParams

	result.Sender = res.Sender
	result.UserID = res.Sender

	ress, err := r.db.ResultFollowOrUnFollow(r.ctx, result)

	if err != nil {
		return db.ResultFollowOrUnFollowRow{}, fmt.Errorf("failed follow :%w", err)
	}

	return ress, nil
}

func (r *followRepository) ResultFollowingList(email string) ([]db.ResultFollowingListRow, error) {

	res_user, err := r.db.GetUsernameOREmail(r.ctx, email)

	if err != nil {
		return nil, fmt.Errorf("failed undefined user: %w", err)
	}

	res, err := r.db.ResultFollowingList(r.ctx, int32(res_user.UserID))

	if err != nil {
		return nil, fmt.Errorf("failed result followlist :%w", err)
	}

	return res, nil
}

func (r *followRepository) ResultFollowersList(email string) ([]db.ResultFollowersListRow, error) {

	res_usr, err := r.db.GetUsernameOREmail(r.ctx, email)

	if err != nil {
		return nil, fmt.Errorf("failed undefined user: %w", err)
	}

	res, err := r.db.ResultFollowersList(r.ctx, int32(res_usr.UserID))

	if err != nil {
		return nil, fmt.Errorf("failed result followersList: %w", err)
	}

	return res, nil
}

func (r *followRepository) SuggestedList(email string) ([]db.SuggestedListRow, error) {
	res_usr, err := r.db.GetUsernameOREmail(r.ctx, email)

	if err != nil {
		return nil, fmt.Errorf("failed undefined user: %w", err)
	}

	res, err := r.db.SuggestedList(r.ctx, int32(res_usr.UserID))

	if err != nil {
		return nil, fmt.Errorf("failed suggedlist :%w", err)
	}

	return res, nil
}
