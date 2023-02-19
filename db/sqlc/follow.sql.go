// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: follow.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const addFollow = `-- name: AddFollow :one
INSERT INTO "follow" ("sender", "receiver", "followStatus", "followOn") VALUES ($1, $2,$3,$4) RETURNING "followID", sender, receiver, "followStatus", "followOn"
`

type AddFollowParams struct {
	Sender       int32     `json:"sender"`
	Receiver     int32     `json:"receiver"`
	FollowStatus string    `json:"followStatus"`
	FollowOn     time.Time `json:"followOn"`
}

func (q *Queries) AddFollow(ctx context.Context, arg AddFollowParams) (Follow, error) {
	row := q.queryRow(ctx, q.addFollowStmt, addFollow,
		arg.Sender,
		arg.Receiver,
		arg.FollowStatus,
		arg.FollowOn,
	)
	var i Follow
	err := row.Scan(
		&i.FollowID,
		&i.Sender,
		&i.Receiver,
		&i.FollowStatus,
		&i.FollowOn,
	)
	return i, err
}

const addFollowerCount = `-- name: AddFollowerCount :one
UPDATE "users" SET "followers"=followers+1 WHERE user_id=$1 RETURNING user_id, "firstName", "lastName", username, email, password, "profileImage", "profileCover", following, followers, bio, country, website
`

func (q *Queries) AddFollowerCount(ctx context.Context, userID int32) (User, error) {
	row := q.queryRow(ctx, q.addFollowerCountStmt, addFollowerCount, userID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.FirstName,
		&i.LastName,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.ProfileImage,
		&i.ProfileCover,
		&i.Following,
		&i.Followers,
		&i.Bio,
		&i.Country,
		&i.Website,
	)
	return i, err
}

const addFollowingCount = `-- name: AddFollowingCount :one
UPDATE "users" SET following=following+1 WHERE "user_id"=$1 RETURNING user_id, "firstName", "lastName", username, email, password, "profileImage", "profileCover", following, followers, bio, country, website
`

func (q *Queries) AddFollowingCount(ctx context.Context, userID int32) (User, error) {
	row := q.queryRow(ctx, q.addFollowingCountStmt, addFollowingCount, userID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.FirstName,
		&i.LastName,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.ProfileImage,
		&i.ProfileCover,
		&i.Following,
		&i.Followers,
		&i.Bio,
		&i.Country,
		&i.Website,
	)
	return i, err
}

const checkFollow = `-- name: CheckFollow :one
SELECT "followID", sender, receiver, "followStatus", "followOn" FROM "follow" WHERE "sender"=$1 AND "receiver"=$2
`

type CheckFollowParams struct {
	Sender   int32 `json:"sender"`
	Receiver int32 `json:"receiver"`
}

func (q *Queries) CheckFollow(ctx context.Context, arg CheckFollowParams) (Follow, error) {
	row := q.queryRow(ctx, q.checkFollowStmt, checkFollow, arg.Sender, arg.Receiver)
	var i Follow
	err := row.Scan(
		&i.FollowID,
		&i.Sender,
		&i.Receiver,
		&i.FollowStatus,
		&i.FollowOn,
	)
	return i, err
}

const removeFollowersCount = `-- name: RemoveFollowersCount :one
UPDATE users SET followers=followers-1 WHERE user_id=$1 RETURNING user_id, "firstName", "lastName", username, email, password, "profileImage", "profileCover", following, followers, bio, country, website
`

func (q *Queries) RemoveFollowersCount(ctx context.Context, userID int32) (User, error) {
	row := q.queryRow(ctx, q.removeFollowersCountStmt, removeFollowersCount, userID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.FirstName,
		&i.LastName,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.ProfileImage,
		&i.ProfileCover,
		&i.Following,
		&i.Followers,
		&i.Bio,
		&i.Country,
		&i.Website,
	)
	return i, err
}

const removeFollowingCount = `-- name: RemoveFollowingCount :one
UPDATE users SET following=following-1 WHERE user_id=$1 RETURNING user_id, "firstName", "lastName", username, email, password, "profileImage", "profileCover", following, followers, bio, country, website
`

func (q *Queries) RemoveFollowingCount(ctx context.Context, userID int32) (User, error) {
	row := q.queryRow(ctx, q.removeFollowingCountStmt, removeFollowingCount, userID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.FirstName,
		&i.LastName,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.ProfileImage,
		&i.ProfileCover,
		&i.Following,
		&i.Followers,
		&i.Bio,
		&i.Country,
		&i.Website,
	)
	return i, err
}

const resultFollowOrUnFollow = `-- name: ResultFollowOrUnFollow :one
SELECT following, followers FROM users LEFT JOIN follow ON sender=$1 AND CASE WHEN receiver=$1 THEN sender=user_id END WHERE user_id=$2
`

type ResultFollowOrUnFollowParams struct {
	Sender int32 `json:"sender"`
	UserID int32 `json:"user_id"`
}

type ResultFollowOrUnFollowRow struct {
	Following int32 `json:"following"`
	Followers int32 `json:"followers"`
}

func (q *Queries) ResultFollowOrUnFollow(ctx context.Context, arg ResultFollowOrUnFollowParams) (ResultFollowOrUnFollowRow, error) {
	row := q.queryRow(ctx, q.resultFollowOrUnFollowStmt, resultFollowOrUnFollow, arg.Sender, arg.UserID)
	var i ResultFollowOrUnFollowRow
	err := row.Scan(&i.Following, &i.Followers)
	return i, err
}

const resultFollowersList = `-- name: ResultFollowersList :many
SELECT user_id, "firstName", "lastName", username, email, password, "profileImage", "profileCover", following, followers, bio, country, website, "followID", sender, receiver, "followStatus", "followOn" FROM users LEFT JOIN follow ON sender=user_id AND CASE WHEN receiver=$1 THEN sender=user_id END WHERE receiver IS NOT NULL ORDER BY followOn DESC
`

type ResultFollowersListRow struct {
	UserID       int32          `json:"user_id"`
	FirstName    string         `json:"firstName"`
	LastName     string         `json:"lastName"`
	Username     string         `json:"username"`
	Email        string         `json:"email"`
	Password     string         `json:"password"`
	ProfileImage string         `json:"profileImage"`
	ProfileCover string         `json:"profileCover"`
	Following    int32          `json:"following"`
	Followers    int32          `json:"followers"`
	Bio          string         `json:"bio"`
	Country      string         `json:"country"`
	Website      string         `json:"website"`
	FollowID     sql.NullInt32  `json:"followID"`
	Sender       sql.NullInt32  `json:"sender"`
	Receiver     sql.NullInt32  `json:"receiver"`
	FollowStatus sql.NullString `json:"followStatus"`
	FollowOn     sql.NullTime   `json:"followOn"`
}

func (q *Queries) ResultFollowersList(ctx context.Context, receiver int32) ([]ResultFollowersListRow, error) {
	rows, err := q.query(ctx, q.resultFollowersListStmt, resultFollowersList, receiver)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ResultFollowersListRow
	for rows.Next() {
		var i ResultFollowersListRow
		if err := rows.Scan(
			&i.UserID,
			&i.FirstName,
			&i.LastName,
			&i.Username,
			&i.Email,
			&i.Password,
			&i.ProfileImage,
			&i.ProfileCover,
			&i.Following,
			&i.Followers,
			&i.Bio,
			&i.Country,
			&i.Website,
			&i.FollowID,
			&i.Sender,
			&i.Receiver,
			&i.FollowStatus,
			&i.FollowOn,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const resultFollowingList = `-- name: ResultFollowingList :many
SELECT user_id, "firstName", "lastName", username, email, password, "profileImage", "profileCover", following, followers, bio, country, website, "followID", sender, receiver, "followStatus", "followOn" FROM "users" LEFT JOIN "follow" ON "receiver"="user_id" AND CASE  WHEN  "sender"=$1 THEN  "receiver"="user_id" END WHERE "sender" IS NOT NULL ORDER BY followOn DESC
`

type ResultFollowingListRow struct {
	UserID       int32          `json:"user_id"`
	FirstName    string         `json:"firstName"`
	LastName     string         `json:"lastName"`
	Username     string         `json:"username"`
	Email        string         `json:"email"`
	Password     string         `json:"password"`
	ProfileImage string         `json:"profileImage"`
	ProfileCover string         `json:"profileCover"`
	Following    int32          `json:"following"`
	Followers    int32          `json:"followers"`
	Bio          string         `json:"bio"`
	Country      string         `json:"country"`
	Website      string         `json:"website"`
	FollowID     sql.NullInt32  `json:"followID"`
	Sender       sql.NullInt32  `json:"sender"`
	Receiver     sql.NullInt32  `json:"receiver"`
	FollowStatus sql.NullString `json:"followStatus"`
	FollowOn     sql.NullTime   `json:"followOn"`
}

func (q *Queries) ResultFollowingList(ctx context.Context, sender int32) ([]ResultFollowingListRow, error) {
	rows, err := q.query(ctx, q.resultFollowingListStmt, resultFollowingList, sender)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ResultFollowingListRow
	for rows.Next() {
		var i ResultFollowingListRow
		if err := rows.Scan(
			&i.UserID,
			&i.FirstName,
			&i.LastName,
			&i.Username,
			&i.Email,
			&i.Password,
			&i.ProfileImage,
			&i.ProfileCover,
			&i.Following,
			&i.Followers,
			&i.Bio,
			&i.Country,
			&i.Website,
			&i.FollowID,
			&i.Sender,
			&i.Receiver,
			&i.FollowStatus,
			&i.FollowOn,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const suggestedList = `-- name: SuggestedList :many
SELECT user_id, "firstName", "lastName", username, email, password, "profileImage", "profileCover", following, followers, bio, country, website, "followID", sender, receiver, "followStatus", "followOn"
FROM users
LEFT JOIN follow ON sender = user_id AND CASE WHEN follow.receiver = $1 THEN sender = user_id END
WHERE user_id != $1 AND follow.receiver IS NULL
INTERSECT
SELECT user_id, "firstName", "lastName", username, email, password, "profileImage", "profileCover", following, followers, bio, country, website, "followID", sender, receiver, "followStatus", "followOn"
FROM users
LEFT JOIN follow ON receiver = user_id AND CASE WHEN follow.sender = $1 THEN receiver = user_id END
WHERE user_id != $1 AND follow.sender IS NULL
ORDER BY followOn DESC
`

type SuggestedListRow struct {
	UserID       int32          `json:"user_id"`
	FirstName    string         `json:"firstName"`
	LastName     string         `json:"lastName"`
	Username     string         `json:"username"`
	Email        string         `json:"email"`
	Password     string         `json:"password"`
	ProfileImage string         `json:"profileImage"`
	ProfileCover string         `json:"profileCover"`
	Following    int32          `json:"following"`
	Followers    int32          `json:"followers"`
	Bio          string         `json:"bio"`
	Country      string         `json:"country"`
	Website      string         `json:"website"`
	FollowID     sql.NullInt32  `json:"followID"`
	Sender       sql.NullInt32  `json:"sender"`
	Receiver     sql.NullInt32  `json:"receiver"`
	FollowStatus sql.NullString `json:"followStatus"`
	FollowOn     sql.NullTime   `json:"followOn"`
}

func (q *Queries) SuggestedList(ctx context.Context, receiver int32) ([]SuggestedListRow, error) {
	rows, err := q.query(ctx, q.suggestedListStmt, suggestedList, receiver)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SuggestedListRow
	for rows.Next() {
		var i SuggestedListRow
		if err := rows.Scan(
			&i.UserID,
			&i.FirstName,
			&i.LastName,
			&i.Username,
			&i.Email,
			&i.Password,
			&i.ProfileImage,
			&i.ProfileCover,
			&i.Following,
			&i.Followers,
			&i.Bio,
			&i.Country,
			&i.Website,
			&i.FollowID,
			&i.Sender,
			&i.Receiver,
			&i.FollowStatus,
			&i.FollowOn,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const unFollow = `-- name: UnFollow :one
DELETE FROM "follow" WHERE "sender"=$1 AND "receiver"=$2 AND "followStatus"=1 RETURNING "followID", sender, receiver, "followStatus", "followOn"
`

type UnFollowParams struct {
	Sender   int32 `json:"sender"`
	Receiver int32 `json:"receiver"`
}

func (q *Queries) UnFollow(ctx context.Context, arg UnFollowParams) (Follow, error) {
	row := q.queryRow(ctx, q.unFollowStmt, unFollow, arg.Sender, arg.Receiver)
	var i Follow
	err := row.Scan(
		&i.FollowID,
		&i.Sender,
		&i.Receiver,
		&i.FollowStatus,
		&i.FollowOn,
	)
	return i, err
}

const whoToFollow = `-- name: WhoToFollow :many
SELECT user_id, "firstName", "lastName", username, email, password, "profileImage", "profileCover", following, followers, bio, country, website FROM "users" WHERE "user_id" != $1 AND "user_id" NOT IN (SELECT "receiver" FROM "follow" WHERE "sender"=$1) ORDER BY random() LIMIT 3
`

func (q *Queries) WhoToFollow(ctx context.Context, userID int32) ([]User, error) {
	rows, err := q.query(ctx, q.whoToFollowStmt, whoToFollow, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.UserID,
			&i.FirstName,
			&i.LastName,
			&i.Username,
			&i.Email,
			&i.Password,
			&i.ProfileImage,
			&i.ProfileCover,
			&i.Following,
			&i.Followers,
			&i.Bio,
			&i.Country,
			&i.Website,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}