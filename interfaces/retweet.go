package interfaces

import (
	db "github.com/renaldyhidayatt/twittersqlc/db/sqlc"
	"github.com/renaldyhidayatt/twittersqlc/dto/request"
)

type IRetweetRepository interface {
	RetweetCount(req request.RetweetCountRequest) (string, error)
	ResetRetweetCount(req request.ResetRetweetRequest) (string, error)

	GetRetweet(tweet_id int) (int64, error)
	CreateRetweet(req request.CreateRetweetRequest) (db.Retweet, error)
}

type IRetweetService interface {
	RetweetCount(req request.RetweetCountRequest) (string, error)
	ResetRetweetCount(req request.ResetRetweetRequest) (string, error)

	GetRetweet(tweet_id int) (int64, error)
	CreateRetweet(req request.CreateRetweetRequest) (db.Retweet, error)
}
