package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	dbConn "github.com/renaldyhidayatt/twittersqlc/db/sqlc"
	"github.com/renaldyhidayatt/twittersqlc/router"
	"github.com/renaldyhidayatt/twittersqlc/utils"
	"github.com/spf13/viper"
)

var (
	db *dbConn.Queries
)

func main() {
	ctx := context.Background()

	err := utils.Viper()

	if err != nil {
		log.Fatal(err.Error())

		panic(err)
	}

	conn, err := utils.Database()

	db = dbConn.New(conn)

	if err != nil {
		log.Fatal(err.Error())

		panic(err)
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	router.NewAuthRouter("/auth", db, ctx, r)
	router.NewCommentRouter("/comment", db, ctx, r)
	router.NewFollowRouter("/follow", db, ctx, r)
	router.NewLikeRouter("/like", db, ctx, r)
	router.NewUserRouter("/user", db, ctx, r)
	router.NewRetweetRouter("/retweet", db, ctx, r)
	router.NewTrendRouter("/trend", db, ctx, r)
	router.NewTweetRouter("/tweet", db, ctx, r)

	serve := &http.Server{
		Addr:           fmt.Sprintf(":%s", viper.GetString("PORT")),
		ReadTimeout:    time.Duration(time.Second) * 60,
		WriteTimeout:   time.Duration(time.Second) * 30,
		IdleTimeout:    time.Duration(time.Second) * 120,
		MaxHeaderBytes: 3145728,
		Handler:        r,
	}

	go func() {
		err := serve.ListenAndServe()

		if err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("Connected to port:", viper.GetString("PORT"))

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	serve.Shutdown(ctx)
	os.Exit(0)
}
