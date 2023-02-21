package services

import (
	db "github.com/renaldyhidayatt/twittersqlc/db/sqlc"
	"github.com/renaldyhidayatt/twittersqlc/dto/request"
	"github.com/renaldyhidayatt/twittersqlc/interfaces"
	"github.com/renaldyhidayatt/twittersqlc/repository"
)

type TweetService = interfaces.ITweetService

type tweetService struct {
	repository repository.TweetRepository
}

func NewTweetService(repository repository.TweetRepository) *tweetService {
	return &tweetService{repository: repository}
}

func (s *tweetService) CreateTweet(req request.CreateTweetRequest) (db.Tweet, error) {
	res, err := s.repository.CreateTweet(req)

	return res, err
}

func (s *tweetService) GetTweet(req request.GetTweetRequest) (db.GetTweetRow, error) {
	res, err := s.repository.GetTweet(req)

	return res, err
}
func (s *tweetService) TweetCounts(tweet_by int) (int64, error) {
	res, err := s.repository.TweetCounts(tweet_by)

	return res, err
}
func (s *tweetService) GetMention(email string) ([]db.User, error) {
	res, err := s.repository.GetMention(email)

	return res, err
}
func (s *tweetService) GetHashTagTweet(hashtag string) ([]db.GetHashTagTweetRow, error) {
	res, err := s.repository.GetHashTagTweet(hashtag)

	return res, err
}
func (s *tweetService) TweetByMe(email string) ([]db.GetTweetByMeRow, error) {
	res, err := s.repository.TweetByMe(email)

	return res, err
}
func (s *tweetService) Tweets(email string) ([]db.GetTweetAllRow, error) {
	res, err := s.repository.Tweets(email)

	return res, err
}
