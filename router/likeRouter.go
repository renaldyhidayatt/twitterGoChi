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

func NewLikeRouter(prefix string, db *db.Queries, ctx context.Context, router *chi.Mux) {
	repository := repository.NewLikeRepository(db, ctx)
	services := services.NewLikeService(repository)
	handler := handler.NewLikesHandler(services)

	router.Route(prefix, func(r chi.Router) {
		r.Use(middlewares.MiddlewareAuthentication)

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello Comment Route"))

		})

		r.Post("/getlikes/{tweet_id}", handler.GetLikes)
		r.Post("/likes", handler.Likes)
	})
}
