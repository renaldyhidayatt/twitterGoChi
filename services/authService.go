package services

import (
	db "github.com/renaldyhidayatt/twittersqlc/db/sqlc"
	"github.com/renaldyhidayatt/twittersqlc/dto/request"
	"github.com/renaldyhidayatt/twittersqlc/interfaces"
)

type AuthService = interfaces.IAuthService

type authService struct {
	repository interfaces.IAuthRepository
}

func NewAuthService(repository interfaces.IAuthRepository) *authService {
	return &authService{repository: repository}
}

func (s *authService) RegisterUser(req request.RegisterRequest) (db.User, error) {
	res, err := s.repository.RegisterUser(req)

	return res, err
}

func (s *authService) Login(req request.LoginRequest) (db.User, error) {
	res, err := s.repository.Login(req)

	return res, err
}
