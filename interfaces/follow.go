package interfaces

import (
	db "github.com/renaldyhidayatt/twittersqlc/db/sqlc"
	"github.com/renaldyhidayatt/twittersqlc/dto/request"
)

type IFollowRepository interface {
	CheckFollow(req request.FollowCheckRequest) (db.Follow, error)
	WhoToFollow(email string) ([]db.User, error)
	Follow(req request.AddFollowRequest) (db.ResultFollowOrUnFollowRow, error)
	UnFollow(req request.UnFollowRequest) (db.ResultFollowOrUnFollowRow, error)

	ResultFollowingList(email string) ([]db.ResultFollowingListRow, error)

	ResultFollowersList(email string) ([]db.ResultFollowersListRow, error)

	SuggestedList(email string) ([]db.SuggestedListRow, error)
}

type IFollowService interface {
	CheckFollow(req request.FollowCheckRequest) (db.Follow, error)
	WhoToFollow(email string) ([]db.User, error)
	Follow(req request.AddFollowRequest) (db.ResultFollowOrUnFollowRow, error)
	UnFollow(req request.UnFollowRequest) (db.ResultFollowOrUnFollowRow, error)

	ResultFollowingList(email string) ([]db.ResultFollowingListRow, error)

	ResultFollowersList(email string) ([]db.ResultFollowersListRow, error)

	SuggestedList(email string) ([]db.SuggestedListRow, error)
}
