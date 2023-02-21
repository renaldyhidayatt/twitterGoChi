package interfaces

import (
	db "github.com/renaldyhidayatt/twittersqlc/db/sqlc"
	"github.com/renaldyhidayatt/twittersqlc/dto/request"
)

type ITweetRepository interface {
	CreateTweet(req request.CreateTweetRequest) (db.Tweet, error)
	GetTweet(req request.GetTweetRequest) (db.GetTweetRow, error)
	TweetCounts(tweet_by int) (int64, error)
	GetMention(email string) ([]db.User, error)
	GetHashTagTweet(hashtag string) ([]db.GetHashTagTweetRow, error)
	TweetByMe(email string) ([]db.GetTweetByMeRow, error)
	Tweets(email string) ([]db.GetTweetAllRow, error)
}

type ITweetService interface {
	CreateTweet(req request.CreateTweetRequest) (db.Tweet, error)
	GetTweet(req request.GetTweetRequest) (db.GetTweetRow, error)
	TweetCounts(tweet_by int) (int64, error)
	GetMention(email string) ([]db.User, error)
	GetHashTagTweet(hashtag string) ([]db.GetHashTagTweetRow, error)
	TweetByMe(email string) ([]db.GetTweetByMeRow, error)
	Tweets(email string) ([]db.GetTweetAllRow, error)
}
