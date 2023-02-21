package interfaces

import (
	db "github.com/renaldyhidayatt/twittersqlc/db/sqlc"
	"github.com/renaldyhidayatt/twittersqlc/dto/request"
)

type IUserRepository interface {
	FindAll() ([]db.User, error)
	GetCurrentUser(username string) (db.User, error)
	UpdateUser(id int, req request.UpdateUserRequest) (db.User, error)
	DeleteUser(user_id int) (bool, error)
}

type IUserService interface {
	FindAll() ([]db.User, error)
	GetCurrentUser(username string) (db.User, error)
	UpdateUser(id int, req request.UpdateUserRequest) (db.User, error)
	DeleteUser(user_id int) (bool, error)
}
