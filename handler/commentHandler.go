package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/renaldyhidayatt/twittersqlc/dto/request"
	"github.com/renaldyhidayatt/twittersqlc/dto/response"
	"github.com/renaldyhidayatt/twittersqlc/services"
)

type commentHandler struct {
	services services.CommentService
}

func NewCommentHandler(services services.CommentService) *commentHandler {
	return &commentHandler{services: services}
}

func (h *commentHandler) RepliestTweet(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value("decoded")
	tokenStr := token.(string)
	var decodedToken map[string]interface{}
	json.Unmarshal([]byte(tokenStr), &decodedToken)
	email, ok := decodedToken["sub"].(string)
	if !ok {
		response.ResponseError(w, http.StatusBadRequest, fmt.Errorf("failed convert"))
	}

	res, err := h.services.RepliesTweet(email)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)
	}
}

func (h *commentHandler) Comment(w http.ResponseWriter, r *http.Request) {
	var commentRequest request.CreateCommentRequest

	err := json.NewDecoder(r.Body).Decode(&commentRequest)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	if err = commentRequest.Validate(); err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	res, err := h.services.Comment(commentRequest)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)
	}
}

func (h *commentHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	var deleteCommentRequest request.DeleteCommentRequest

	err := json.NewDecoder(r.Body).Decode(&deleteCommentRequest)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	if err = deleteCommentRequest.Validate(); err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return

	}
	res, err := h.services.DeleteComment(deleteCommentRequest)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil menghapus data", res, http.StatusOK)
	}
}
