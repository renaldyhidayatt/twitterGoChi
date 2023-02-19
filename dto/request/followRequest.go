package request

type AddFollowRequest struct {
	Sender   int32 `json:"sender"`
	Receiver int32 `json:"receiver"`
}
