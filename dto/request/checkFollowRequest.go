package request

type FollowCheckRequest struct {
	Following int `json:"following" validate:"required"`
	UserID    int `json:"user_id" validate:"required"`
}
