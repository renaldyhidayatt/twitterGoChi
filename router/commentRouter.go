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

func NewCommentRouter(prefix string, db *db.Queries, ctx context.Context, router *chi.Mux) {
	repository := repository.NewCommentRepository(db, ctx)
	services := services.NewCommentService(repository)
	handler := handler.NewCommentHandler(services)

	router.Route(prefix, func(r chi.Router) {
		r.Use(middlewares.MiddlewareAuthentication)

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello Comment Route"))

		})

		r.Post("/repliestweet", handler.RepliestTweet)
		r.Post("/comment", handler.Comment)
		r.Delete("/deletecomment", handler.DeleteComment)
	})
}
