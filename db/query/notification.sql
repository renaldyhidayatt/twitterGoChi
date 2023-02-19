-- name: CreateNotification :one
INSERT INTO notification("notificationFor", "notificationFrom", type, status, target,"notificationCount", "notificationOn") VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING *;

-- name: DeleteNotification :exec
DELETE FROM notification WHERE "notificationFor"=$1 AND "notificationFrom"=$2 AND type=$3 AND target = $4;

