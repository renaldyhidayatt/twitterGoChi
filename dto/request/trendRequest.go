package request

type TrendRequest struct {
	Hashtag string `json:"hashtag"`
	TweetId int32  `json:"tweetId"`
	UserID  int32  `json:"user_id"`
}
