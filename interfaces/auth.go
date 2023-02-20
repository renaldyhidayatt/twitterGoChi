package interfaces

import (
	db "github.com/renaldyhidayatt/twittersqlc/db/sqlc"
	"github.com/renaldyhidayatt/twittersqlc/dto/request"
)

type IAuthRepository interface {
	RegisterUser(req request.RegisterRequest) (db.User, error)
	Login(req request.LoginRequest) (db.User, error)
}

type IAuthService interface {
	RegisterUser(req request.RegisterRequest) (db.User, error)
	Login(req request.LoginRequest) (db.User, error)
}
