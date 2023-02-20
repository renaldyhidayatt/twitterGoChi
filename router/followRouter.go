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

func NewFollowRouter(prefix string, db *db.Queries, ctx context.Context, router *chi.Mux) {
	repository := repository.NewFollowRepository(db, ctx)
	services := services.NewFollowService(repository)
	handler := handler.NewFollowHandler(services)

	router.Route(prefix, func(r chi.Router) {
		r.Use(middlewares.MiddlewareAuthentication)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello Follow Router"))
		})

		r.Post("/checkfollow", handler.CheckFollow)
		r.Post("/whotofollow", handler.WhoToFollow)
		r.Post("/follow", handler.Follow)
		r.Post("/unfollow", handler.Unfollow)
		r.Post("/resultfollowinglist", handler.ResultFollowingList)
		r.Post("/resultfollowerslist", handler.ResultFollowersList)
		r.Post("/suggestedlist", handler.SuggestedList)

	})
}
