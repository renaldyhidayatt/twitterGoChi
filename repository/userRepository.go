package repository

import (
	"context"
	"fmt"

	db "github.com/renaldyhidayatt/twittersqlc/db/sqlc"
	"github.com/renaldyhidayatt/twittersqlc/dto/request"
	"github.com/renaldyhidayatt/twittersqlc/security"
)

type userRepository struct {
	db  *db.Queries
	ctx context.Context
}

func NewUserRepository(db *db.Queries, ctx context.Context) *userRepository {
	return &userRepository{db: db, ctx: ctx}
}

func (r *userRepository) FindAll() ([]db.User, error) {
	res, err := r.db.GetUsers(r.ctx)

	if err != nil {
		return nil, fmt.Errorf("failed results user: %w", err)
	}

	return res, nil
}

func (r *userRepository) GetCurrentUser(id int) (db.User, error) {
	res, err := r.db.GetCurrentUser(r.ctx, int32(id))

	if err != nil {
		return db.User{}, fmt.Errorf("failed result user: %w", err)
	}

	return res, nil
}

func (r *userRepository) GetUsername(username string) (db.User, error) {
	res, err := r.db.GetUsernameOREmail(r.ctx, username)

	if err != nil {
		return db.User{}, fmt.Errorf("failed result user: %w", err)
	}

	return res, nil
}

func (r *userRepository) GetHashTagTweet(hashtag string) ([]db.GetHashTagTweetRow, error) {
	res, err := r.db.GetHashTagTweet(r.ctx, hashtag)

	if err != nil {
		return nil, fmt.Errorf("failed result user hashtag: %w", err)
	}

	return res, nil
}

func (r *userRepository) GetMention(username string) ([]db.User, error) {
	res, err := r.db.GetMention(r.ctx, username)

	if err != nil {
		return nil, fmt.Errorf("failed result user mention: %w", err)
	}

	return res, nil
}

func (r *userRepository) UpdateUser(id int, req request.UpdateUserRequest) (db.User, error) {
	var updateUser db.UpdateUserParams

	res, err := r.db.GetCurrentUser(r.ctx, int32(id))

	if err != nil {
		return db.User{}, fmt.Errorf("failed result user: %w", err)
	}

	updateUser.FirstName = req.Firstname
	updateUser.LastName = req.Lastname
	updateUser.Username = res.Username
	updateUser.Email = req.Email
	updateUser.Password = security.HashPassword(req.Password)
	updateUser.Username_2 = req.Username_2
	updateUser.ProfileImage = req.Profileimage
	updateUser.ProfileCover = req.Profilecover
	updateUser.Bio = req.Bio
	updateUser.Country = req.Country
	updateUser.Website = req.Website

	res, err = r.db.UpdateUser(r.ctx, updateUser)

	if err != nil {
		return db.User{}, fmt.Errorf("failed update user: %w", err)
	}

	return res, nil
}

func (r *userRepository) DeleteUser(user_id int) (bool, error) {
	res, err := r.db.GetCurrentUser(r.ctx, int32(user_id))

	if err != nil {
		return false, fmt.Errorf("failed result user: %w", err)
	}

	err = r.db.DeleteUser(r.ctx, res.Username)

	if err != nil {
		return false, fmt.Errorf("failed delete user: %w", err)
	}

	return true, nil
}
