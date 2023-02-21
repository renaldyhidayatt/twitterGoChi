package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/renaldyhidayatt/twittersqlc/dto/request"
	"github.com/renaldyhidayatt/twittersqlc/dto/response"
	"github.com/renaldyhidayatt/twittersqlc/services"
)

type trendHandler struct {
	services services.TrendService
}

func NewTrendHandler(services services.TrendService) *trendHandler {
	return &trendHandler{services: services}

}

func (h *trendHandler) GetTrend(w http.ResponseWriter, r *http.Request) {
	res, err := h.services.GetTrend()

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)
	}
}

func (h *trendHandler) GetTrendHash(w http.ResponseWriter, r *http.Request) {
	hash := chi.URLParam(r, "hashtag")

	res, err := h.services.GetTrendByHash(hash)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)
	}
}

func (h *trendHandler) CreateTrend(w http.ResponseWriter, r *http.Request) {
	var trend request.TrendRequest

	err := json.NewDecoder(r.Body).Decode(&trend)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	if err = trend.Validate(); err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}
}
