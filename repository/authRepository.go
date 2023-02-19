package repository

import (
	"context"
	"fmt"

	db "github.com/renaldyhidayatt/twittersqlc/db/sqlc"
	"github.com/renaldyhidayatt/twittersqlc/dto/request"
	"github.com/renaldyhidayatt/twittersqlc/security"
)

type authRepository struct {
	db  *db.Queries
	ctx context.Context
}

func NewAuthRepository(db *db.Queries, ctx context.Context) *authRepository {
	return &authRepository{db: db, ctx: ctx}
}

func (r *authRepository) RegisterUser(req request.RegisterRequest) (db.User, error) {
	var createUser db.CreateUserParams

	createUser.FirstName = req.Firstname
	createUser.LastName = req.Lastname
	createUser.Email = req.Email
	createUser.Password = security.HashPassword(req.Password)

	res, err := r.db.CreateUser(r.ctx, createUser)

	if err != nil {
		return db.User{}, fmt.Errorf("failed create user :%w", err)
	}

	return res, nil
}

func (r *authRepository) Login(req request.LoginRequest) (db.User, error) {

	res, err := r.db.GetUsernameOREmail(r.ctx, req.Username)

	if err != nil {
		return db.User{}, fmt.Errorf("failed username not found :%w", err)
	}

	checkPassword := security.VerifyPassword(res.Password, req.Password)

	if checkPassword != nil {
		return db.User{}, fmt.Errorf("failed checkhash password: %w", err)
	}

	return res, nil

}
