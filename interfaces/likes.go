package interfaces

import "github.com/renaldyhidayatt/twittersqlc/dto/request"

type ILikesRepository interface {
	GetLikes(tweet_id int) (int64, error)
	Likes(req request.LikesRequest) (map[string]int, error)
}

type ILikesService interface {
	GetLikes(tweet_id int) (int64, error)
	Likes(req request.LikesRequest) (map[string]int, error)
}
