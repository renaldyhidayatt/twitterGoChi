-- name: GetTrends :many
SELECT *,COUNT("tweet_id") AS "tweetsCount" FROM trends t LEFT JOIN tweet p ON p.tweet_id=t.tweetId AND status LIKE CONCAT("%#","hashtag","%") GROUP BY "hashtag" ORDER BY "tweetsCount" DESC LIMIT 3;

-- name: GetTrendByHash :many
SELECT DISTINCT "hashtag" FROM "trends" WHERE "hashtag" LIKE $1 LIMIT 5;

-- name: CreateTrend :one
INSERT INTO "trends" ("hashtag","tweetId","user_id","createdOn") VALUES ($1,$2,$3,$4) RETURNING *;

