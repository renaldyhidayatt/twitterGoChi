package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/renaldyhidayatt/twittersqlc/dto/request"
	"github.com/renaldyhidayatt/twittersqlc/dto/response"
	"github.com/renaldyhidayatt/twittersqlc/services"
)

type tweetHandler struct {
	services services.TweetService
}

func NewTweetHandler(services services.TweetService) *tweetHandler {
	return &tweetHandler{services: services}
}

func (h *tweetHandler) Tweets(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value("decoded")
	tokenStr := token.(string)
	var decodedToken map[string]interface{}
	json.Unmarshal([]byte(tokenStr), &decodedToken)
	email, ok := decodedToken["sub"].(string)
	if !ok {
		response.ResponseError(w, http.StatusBadRequest, fmt.Errorf("failed convert"))
	}

	res, err := h.services.Tweets(email)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)
	}
}

func (h *tweetHandler) TweetByMe(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value("decoded")
	tokenStr := token.(string)
	var decodedToken map[string]interface{}
	json.Unmarshal([]byte(tokenStr), &decodedToken)
	email, ok := decodedToken["sub"].(string)
	if !ok {
		response.ResponseError(w, http.StatusBadRequest, fmt.Errorf("failed convert"))
	}

	res, err := h.services.TweetByMe(email)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)
	}
}

func (h *tweetHandler) GetHashTagTweet(w http.ResponseWriter, r *http.Request) {
	hashtag := chi.URLParam(r, "hashtag")

	res, err := h.services.GetHashTagTweet(hashtag)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)
	}
}

func (h *tweetHandler) GetMention(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")

	res, err := h.services.GetMention(email)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)
	}
}

func (h *tweetHandler) TweetCounts(w http.ResponseWriter, r *http.Request) {
	tweet_by := chi.URLParam(r, "tweet_by")

	tweet_int, err := strconv.Atoi(tweet_by)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	res, err := h.services.TweetCounts(tweet_int)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)
	}
}

func (h *tweetHandler) GetTweet(w http.ResponseWriter, r *http.Request) {
	var request request.GetTweetRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = request.Validate(); err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	res, err := h.services.GetTweet(request)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)
	}
}

func (h *tweetHandler) CreateTweet(w http.ResponseWriter, r *http.Request) {
	var createrequest request.CreateTweetRequest

	err := json.NewDecoder(r.Body).Decode(&createrequest)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = createrequest.Validate(); err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	res, err := h.services.CreateTweet(createrequest)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil membuat data", res, http.StatusOK)
	}
}
