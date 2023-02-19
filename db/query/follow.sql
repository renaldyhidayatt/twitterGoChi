-- name: CheckFollow :one
SELECT * FROM "follow" WHERE "sender"=$1 AND "receiver"=$2;

-- name: WhoToFollow :many
SELECT * FROM "users" WHERE "user_id" != $1 AND "user_id" NOT IN (SELECT "receiver" FROM "follow" WHERE "sender"=$1) ORDER BY random() LIMIT 3;

-- name: AddFollow :one
INSERT INTO "follow" ("sender", "receiver", "followStatus", "followOn") VALUES ($1, $2,$3,$4) RETURNING *;

-- name: UnFollow :one
DELETE FROM "follow" WHERE "sender"=$1 AND "receiver"=$2 AND "followStatus"=1 RETURNING *;


-- name: AddFollowingCount :one
UPDATE "users" SET following=following+1 WHERE "user_id"=$1 RETURNING *;


-- name: AddFollowerCount :one
UPDATE "users" SET "followers"=followers+1 WHERE user_id=$1 RETURNING *;

-- name: RemoveFollowingCount :one
UPDATE users SET following=following-1 WHERE user_id=$1 RETURNING *;

-- name: RemoveFollowersCount :one
UPDATE users SET followers=followers-1 WHERE user_id=$1 RETURNING *;

-- name: ResultFollowOrUnFollow :one
SELECT following, followers FROM users LEFT JOIN follow ON sender=$1 AND CASE WHEN receiver=$1 THEN sender=user_id END WHERE user_id=$2;

-- name: ResultFollowingList :many
SELECT * FROM "users" LEFT JOIN "follow" ON "receiver"="user_id" AND CASE  WHEN  "sender"=$1 THEN  "receiver"="user_id" END WHERE "sender" IS NOT NULL ORDER BY followOn DESC;

-- name: ResultFollowersList :many
SELECT * FROM users LEFT JOIN follow ON sender=user_id AND CASE WHEN receiver=$1 THEN sender=user_id END WHERE receiver IS NOT NULL ORDER BY followOn DESC;

-- name: SuggestedList :many
SELECT *
FROM users
LEFT JOIN follow ON sender = user_id AND CASE WHEN follow.receiver = $1 THEN sender = user_id END
WHERE user_id != $1 AND follow.receiver IS NULL
INTERSECT
SELECT *
FROM users
LEFT JOIN follow ON receiver = user_id AND CASE WHEN follow.sender = $1 THEN receiver = user_id END
WHERE user_id != $1 AND follow.sender IS NULL
ORDER BY followOn DESC;
