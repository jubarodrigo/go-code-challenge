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

	"go-code-challenge/apiserver"
	"go-code-challenge/datastore"
	"go-code-challenge/datastore/files/repositories/users_actions"
	"go-code-challenge/internal"
	services2 "go-code-challenge/internal/actions/services"
	"go-code-challenge/internal/users/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/docgen"
	"github.com/jessevdk/go-flags"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
)

type Args struct {
	Address  string `short:"a" long:"address" description:"The address to listen on for HTTP requests" default:"localhost"`
	Port     int    `short:"p" long:"port" description:"The port to listen on for HTTP requests" default:"3333"`
	Routes   bool   `short:"r" long:"routes" description:"Generate router documentation"`
	Database bool   `short:"d" long:"database" description:"Use a database"`
}

type CodeChallenge struct {
	args          Args
	mode          string
	userService   internal.UserServiceInterface
	actionService internal.ActionServiceInterface
	dataRepo      datastore.DatasJsonRepositoryInterface
}

func NewApp() *CodeChallenge {
	return &CodeChallenge{}
}

type AppBuilder interface {
	setConfig()
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
	sbd.builder.setRepositories()
	sbd.builder.setServices()
	sbd.builder.setWebServer()
}

func (cch *CodeChallenge) setConfig() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	parser := flags.NewParser(&cch.args, flags.Default)
	_, err := parser.Parse()
	if err != nil {
		log.Err(err)
		return
	}
}

func (cch *CodeChallenge) setWebServer() {
	log.Info().Msg("Starting Web Server...")

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cch.args.Address, cch.args.Port),
		Handler: cch.setupRouter(),
	}

	go func() {
		log.Info().Msgf("Listening on %s:%d", cch.args.Address, cch.args.Port)
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

func (cch *CodeChallenge) setupRouter() http.Handler {
	router := chi.NewRouter()
	webServer := apiserver.NewServer(cch.actionService, cch.userService)
	webServer.SetupRoutes(router)

	if cch.args.Routes {
		log.Info().Msg(docgen.JSONRoutesDoc(router))
	}

	return router
}

func (cch *CodeChallenge) setRepositories() {
	cch.dataRepo = users_actions.NewJSONRepository()
}

func (cch *CodeChallenge) setServices() {
	cch.userService = services.NewUserService(cch.dataRepo)
	cch.actionService = services2.NewActionService(cch.dataRepo)
}
