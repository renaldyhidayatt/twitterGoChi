-- name: GetLikes :one
SELECT  count(*) as "count" FROM "likes" WHERE "likeOn"=$1;

-- name: WasLikeBy :one
SELECT  * FROM "likes" WHERE "likeOn"=$1 AND "likeBy" = $2;

-- name: DeleteLike :exec
DELETE FROM likes WHERE "likeBy" = $1 AND "likeOn" = $2;

-- name: CreateLike :one
INSERT INTO likes("likeBy", "likeOn") VALUES ($1, $2) RETURNING *;