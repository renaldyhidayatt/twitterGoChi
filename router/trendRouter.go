package router

import (
	"context"

	"github.com/go-chi/chi/v5"
	db "github.com/renaldyhidayatt/twittersqlc/db/sqlc"
	"github.com/renaldyhidayatt/twittersqlc/handler"
	"github.com/renaldyhidayatt/twittersqlc/middlewares"
	"github.com/renaldyhidayatt/twittersqlc/repository"
	"github.com/renaldyhidayatt/twittersqlc/services"
)

func NewTrendRouter(prefix string, db *db.Queries, ctx context.Context, router *chi.Mux) {
	repository := repository.NewTrendRepository(db, ctx)
	services := services.NewTrendService(repository)
	handler := handler.NewTrendHandler(services)

	router.Route(prefix, func(r chi.Router) {
		r.Use(middlewares.MiddlewareAuthentication)

		r.Post("/gettrend", handler.GetTrend)
		r.Post("/{hashtag}", handler.GetTrendHash)
		r.Post("/create", handler.CreateTrend)
	})
}
