package engine

import (
	"context"
	"github.com/deltrinos/rss-api/app/conf"
	"github.com/deltrinos/rss-api/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Engine struct {
	addr    string
	Handler *gin.Engine
	Queries interfaces.IStorageQueries
}

func NewEngine(addr string) *Engine {
	if conf.Env.IsProduction {
		gin.SetMode(gin.ReleaseMode)
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	return &Engine{
		addr:    addr,
		Handler: gin.Default(),
	}
}

func (e *Engine) Start() {
	srv := http.Server{
		Addr:         e.addr,
		Handler:      e.Handler,
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Error().Msgf("server listenAndServe: %v", err)
		}
	}()
	<-done

	log.Info().Msgf("server shutting down...")
	err := srv.Shutdown(context.Background())
	if err != nil {
		log.Fatal().Msgf("server shutdown: %v", err)
	}
	log.Info().Msgf("exiting.")
}
