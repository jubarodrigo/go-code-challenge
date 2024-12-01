package cmd

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/docgen"
	"github.com/jessevdk/go-flags"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"

	"go-code-challenge/apiserver"
)

type Args struct {
	Address  string `short:"a" long:"address" description:"The address to listen on for HTTP requests" default:"localhost"`
	Port     int    `short:"p" long:"port" description:"The port to listen on for HTTP requests" default:"3333"`
	Routes   bool   `short:"r" long:"routes" description:"Generate router documentation"`
	Database bool   `short:"d" long:"database" description:"Use a database"`
}

type CodeChallenge struct {
	args Args
	mode string
}

func NewApp() *CodeChallenge {
	return &CodeChallenge{}
}

type AppBuilder interface {
	setConfig()
	setDatabase()
	setRepositories()
	setServices()
	setWebServer()
}

type BuildDirector struct {
	builder AppBuilder
}

func NewAppBuilder(ab AppBuilder) *BuildDirector {
	return &BuildDirector{
		builder: ab,
	}
}

func (sbd *BuildDirector) BuildStarship() {
	sbd.builder.setConfig()
	sbd.builder.setDatabase()
	sbd.builder.setRepositories()
	sbd.builder.setServices()
	sbd.builder.setWebServer()
}

func (star *CodeChallenge) setConfig() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	parser := flags.NewParser(&star.args, flags.Default)
	_, err := parser.Parse()
	if err != nil {
		log.Err(err)
		return
	}
}

func (star *CodeChallenge) setWebServer() {
	log.Info().Msg("Starting Web Server...")

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", star.args.Address, star.args.Port),
		Handler: star.setupRouter(),
	}

	go func() {
		log.Info().Msgf("Listening on %s:%d", star.args.Address, star.args.Port)
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msg("Server error")
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("Server shutdown failed")
	}

	log.Info().Msg("Server gracefully stopped")
}

func (star *CodeChallenge) setupRouter() http.Handler {
	router := chi.NewRouter()
	webServer := apiserver.NewServer()
	webServer.SetupRoutes(router)

	if star.args.Routes {
		log.Info().Msg(docgen.JSONRoutesDoc(router))
	}

	return router
}

func (star *CodeChallenge) setDatabase() {}

func (star *CodeChallenge) setRepositories() {
}

func (star *CodeChallenge) setServices() {
}
