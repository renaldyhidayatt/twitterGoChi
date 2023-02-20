package router

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	db "github.com/renaldyhidayatt/twittersqlc/db/sqlc"
	"github.com/renaldyhidayatt/twittersqlc/handler"
	"github.com/renaldyhidayatt/twittersqlc/middlewares"
	"github.com/renaldyhidayatt/twittersqlc/repository"
	"github.com/renaldyhidayatt/twittersqlc/services"
)

func NewRetweetRouter(prefix string, db *db.Queries, ctx context.Context, router *chi.Mux) {
	repository := repository.NewRetweetRepository(db, ctx)
	services := services.NewRetweetService(repository)
	handler := handler.NewRetweetHandler(services)

	router.Route(prefix, func(r chi.Router) {
		r.Use(middlewares.MiddlewareAuthentication)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello Retweet Router"))
		})

		r.Post("/retweetcount", handler.RetweetCount)
		r.Post("/resetretweetcount", handler.ResetRetweetCount)
		r.Post("/getretweet/{tweet_id}", handler.GetRetweet)
	})
}
