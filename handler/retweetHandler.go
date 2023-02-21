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

type retweetHandler struct {
	services services.RetweetService
}

func NewRetweetHandler(services services.RetweetService) *retweetHandler {
	return &retweetHandler{services: services}
}

func (h *retweetHandler) RetweetCount(w http.ResponseWriter, r *http.Request) {
	var retwetcount request.RetweetCountRequest

	err := json.NewDecoder(r.Body).Decode(&retwetcount)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	if err = retwetcount.Validate(); err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	res, err := h.services.RetweetCount(retwetcount)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)
	}
}

func (h *retweetHandler) ResetRetweetCount(w http.ResponseWriter, r *http.Request) {
	var resetCount request.ResetRetweetRequest

	err := json.NewDecoder(r.Body).Decode(&resetCount)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	if err = resetCount.Validate(); err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	res, err := h.services.ResetRetweetCount(resetCount)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)
	}
}

func (h *retweetHandler) GetRetweet(w http.ResponseWriter, r *http.Request) {
	tweet_id := chi.URLParam(r, "tweet_id")

	tweet_id_int, err := strconv.Atoi(tweet_id)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	res, err := h.services.GetRetweet(tweet_id_int)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	response.ResponseMessage(w, "Berhasil Retweet ", res, http.StatusOK)
}

func (h *retweetHandler) CreateRetweet(w http.ResponseWriter, r *http.Request) {
	var retweet request.CreateRetweetRequest

	err := json.NewDecoder(r.Body).Decode(&retweet)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	if err = retweet.Validate(); err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	res, err := h.services.CreateRetweet(retweet)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)
	}
}
