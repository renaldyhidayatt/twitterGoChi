-- name: GetUsers :many
SELECT * FROM users;

-- name: GetCurrentUser :one
SELECT * FROM users WHERE user_id = $1;

-- name: GetUsernameOREmail :one
SELECT * FROM users WHERE username = $1 OR email = $1;

-- name: CreateUser :one
INSERT INTO users ("firstName","lastName","username","email","password","profileImage","profileCover") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING *;

-- name: UpdateUser :one
UPDATE users SET "firstName" = $2, "lastName" = $3, "username" = $4, "email" = $5, "password" = $6, "profileImage" = $7, "profileCover" = $8, "bio" = $9, "country" = $10, "website" = $11  WHERE username = $1 RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE username = $1;