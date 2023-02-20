package interfaces

import "github.com/renaldyhidayatt/twittersqlc/dto/request"

type IRetweetRepository interface {
	RetweetCount(req request.RetweetCountRequest) (string, error)
	ResetRetweetCount(req request.ResetRetweetRequest) (string, error)

	GetRetweet(tweet_id int) (int64, error)
}

type IRetweetService interface {
	RetweetCount(req request.RetweetCountRequest) (string, error)
	ResetRetweetCount(req request.ResetRetweetRequest) (string, error)

	GetRetweet(tweet_id int) (int64, error)
}
