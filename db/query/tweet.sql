-- name: GetTweetAll :many
SELECT * FROM tweet t LEFT JOIN users u ON t."tweetBy"=u.user_id WHERE t."tweetBy"=$1 UNION SELECT * FROM tweet t LEFT JOIN users u ON t."tweetBy"=u.user_id WHERE t."tweetBy" IN (SELECT follow.receiver FROM follow WHERE follow.sender=$1) ORDER BY postedOn;

-- name: GetTweetByMe :many
SELECT * FROM tweet, users WHERE "tweetBy"=user_id AND user_id=$1 ORDER BY postedOn DESC;

-- name: GetHashTagTweet :many
SELECT * FROM users u LEFT JOIN tweet p ON p.tweetBy=u.user_id INNER JOIN trends t ON p.tweet_id=t.tweetId WHERE hashtag=$1 ORDER BY postedOn DESC;

-- name: GetMention :many
SELECT  * FROM "users" WHERE "username" LIKE $1 OR "firstName" LIKE $1 OR "lastName" LIKE $1 LIMIT 5;

-- name: TweetCounts :one
SELECT  count(*) as "count" FROM "tweet" WHERE "tweetBy"=$1;

-- name: GetTweet :one
SELECT * FROM "tweet" LEFT JOIN "users" ON users.user_id=tweet."tweetBy" WHERE tweet.tweet_id=$1 AND tweet."tweetBy"=$2;

-- name: CreateTweet :one
INSERT INTO "tweet" ("status", "tweetBy", "tweetImage", "postedOn") VALUES ($1,$2,$3,$4) RETURNING *;

