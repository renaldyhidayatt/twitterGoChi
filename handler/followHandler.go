package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/renaldyhidayatt/twittersqlc/dto/request"
	"github.com/renaldyhidayatt/twittersqlc/dto/response"
	"github.com/renaldyhidayatt/twittersqlc/services"
)

type followHandler struct {
	services services.FollowService
}

func NewFollowHandler(services services.FollowService) *followHandler {
	return &followHandler{services: services}
}

func (h *followHandler) CheckFollow(w http.ResponseWriter, r *http.Request) {
	var followCheck request.FollowCheckRequest

	err := json.NewDecoder(r.Body).Decode(&followCheck)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	res, err := h.services.CheckFollow(followCheck)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil Check follow", res, http.StatusOK)
	}
}

func (h *followHandler) WhoToFollow(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value("decoded")
	tokenStr := token.(string)
	var decodedToken map[string]interface{}
	json.Unmarshal([]byte(tokenStr), &decodedToken)
	email, ok := decodedToken["sub"].(string)
	if !ok {
		response.ResponseError(w, http.StatusBadRequest, fmt.Errorf("failed convert"))
	}
	res, err := h.services.WhoToFollow(email)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)
}

func (h *followHandler) Follow(w http.ResponseWriter, r *http.Request) {
	var followRequest request.AddFollowRequest

	err := json.NewDecoder(r.Body).Decode(&followRequest)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	if err = followRequest.Validate(); err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	res, err := h.services.Follow(followRequest)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil follow", res, http.StatusOK)
	}

}

func (h *followHandler) Unfollow(w http.ResponseWriter, r *http.Request) {
	var unfollowRequest request.UnFollowRequest

	err := json.NewDecoder(r.Body).Decode(&unfollowRequest)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	if err = unfollowRequest.Validate(); err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}
	res, err := h.services.UnFollow(unfollowRequest)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil follow", res, http.StatusOK)
	}
}

func (h *followHandler) ResultFollowingList(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value("decoded")
	tokenStr := token.(string)
	var decodedToken map[string]interface{}
	json.Unmarshal([]byte(tokenStr), &decodedToken)
	email, ok := decodedToken["sub"].(string)
	if !ok {
		response.ResponseError(w, http.StatusBadRequest, fmt.Errorf("failed convert"))
	}

	res, err := h.services.ResultFollowingList(email)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)
}

func (h *followHandler) ResultFollowersList(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value("decoded")
	tokenStr := token.(string)
	var decodedToken map[string]interface{}
	json.Unmarshal([]byte(tokenStr), &decodedToken)
	email, ok := decodedToken["sub"].(string)
	if !ok {
		response.ResponseError(w, http.StatusBadRequest, fmt.Errorf("failed convert"))
	}

	res, err := h.services.ResultFollowersList(email)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)
}

func (h *followHandler) SuggestedList(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value("decoded")
	tokenStr := token.(string)
	var decodedToken map[string]interface{}
	json.Unmarshal([]byte(tokenStr), &decodedToken)
	email, ok := decodedToken["sub"].(string)
	if !ok {
		response.ResponseError(w, http.StatusBadRequest, fmt.Errorf("failed convert"))
	}

	res, err := h.services.SuggestedList(email)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)
}
