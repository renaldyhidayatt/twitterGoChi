-- name: WasRetweetBy :one
SELECT  * FROM "retweet" WHERE "retweetFrom"=$1 AND "retweetBy"=$2;

-- name: CheckRetweet :one
SELECT  * FROM "retweet" WHERE "retweetFrom"=$1 AND "retweetBy"=$2;

-- name: GetRetweet :one
SELECT  count(*) as "count" FROM "retweet" WHERE "retweetFrom"=$1;

-- name: DeleteRetweet :exec
DELETE FROM retweet WHERE "retweetBy" = $1 AND "retweetFrom" = $2;

-- name: CreateRetweet :one
INSERT INTO retweet ("retweetBy", "retweetFrom") VALUES ($1, $2) RETURNING *;