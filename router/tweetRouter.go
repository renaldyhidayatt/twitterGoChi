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

func NewTweetRouter(prefix string, db *db.Queries, ctx context.Context, router *chi.Mux) {
	repository := repository.NewTweetRepository(db, ctx)
	services := services.NewTweetService(repository)
	handler := handler.NewTweetHandler(services)

	router.Route(prefix, func(r chi.Router) {
		r.Use(middlewares.MiddlewareAuthentication)

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello Comment Route"))

		})
		r.Post("/tweets", handler.Tweets)
		r.Post("/tweetbyme", handler.TweetByMe)
		r.Post("/{hashtag}", handler.GetHashTagTweet)
		r.Post("/{email}", handler.GetMention)
		r.Post("/tweetcounts", handler.TweetCounts)
		r.Post("/getweet", handler.GetTweet)
		r.Post("/create", handler.CreateTweet)

	})
}
