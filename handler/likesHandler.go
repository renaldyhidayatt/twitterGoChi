package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/renaldyhidayatt/twittersqlc/dto/request"
	"github.com/renaldyhidayatt/twittersqlc/dto/response"
	"github.com/renaldyhidayatt/twittersqlc/services"
)

type likesHandler struct {
	services services.LikesService
}

func NewLikesHandler(services services.LikesService) *likesHandler {
	return &likesHandler{services: services}
}

func (h *likesHandler) GetLikes(w http.ResponseWriter, r *http.Request) {
	tweet_id := chi.URLParam(r, "tweet_id")

	tweet_id_int, err := strconv.Atoi(tweet_id)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	res, err := h.services.GetLikes(tweet_id_int)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	response.ResponseMessage(w, "Berhasil mendapatkan data ", res, http.StatusOK)
}

func (h *likesHandler) Likes(w http.ResponseWriter, r *http.Request) {
	var likesrequest request.LikesRequest

	err := json.NewDecoder(r.Body).Decode(&likesrequest)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	if err = likesrequest.Validate(); err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	res, err := h.services.Likes(likesrequest)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)
	}

}
