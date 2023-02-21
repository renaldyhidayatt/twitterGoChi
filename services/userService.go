package services

import (
	db "github.com/renaldyhidayatt/twittersqlc/db/sqlc"
	"github.com/renaldyhidayatt/twittersqlc/dto/request"
	"github.com/renaldyhidayatt/twittersqlc/interfaces"
	"github.com/renaldyhidayatt/twittersqlc/repository"
)

type UserService = interfaces.IUserService

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{repository: repository}
}

func (s *userService) FindAll() ([]db.User, error) {
	res, err := s.repository.FindAll()

	return res, err
}

func (s *userService) GetCurrentUser(username string) (db.User, error) {
	res, err := s.repository.GetCurrentUser(username)

	return res, err
}

func (s *userService) UpdateUser(id int, req request.UpdateUserRequest) (db.User, error) {
	res, err := s.repository.UpdateUser(id, req)

	return res, err
}

func (s *userService) DeleteUser(user_id int) (bool, error) {
	res, err := s.repository.DeleteUser(user_id)

	return res, err
}
