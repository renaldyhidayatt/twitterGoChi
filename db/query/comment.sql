-- name: RepliesTweets :many
SELECT * FROM "comment" LEFT JOIN "users" ON "commentBy"="user_id" WHERE "commentBy"= $1 ORDER BY commentAt DESC;

-- name: GetComment :one
SELECT * FROM "comment" LEFT JOIN "users" ON users.user_id=comment."commentBy" WHERE comment."commentID"=$1 AND comment."commentBy"=$2;

-- name: GetCommentcount :one
SELECT  count(*) as "count" FROM "comment" WHERE "commentOn"=$1;

-- name: DeleteComment :exec
DELETE FROM comment WHERE "commentID" = $1;

-- name: WasCommentBy :one
SELECT  * FROM "comment" WHERE "commentOn"=$1 AND "commentBy" = $2;

-- name: CreateComment :one
INSERT INTO comment ("commentBy", "commentOn", "comment","commentAt") VALUES ($1, $2, $3,$4) RETURNING *;