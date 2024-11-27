package cmd

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/docgen"
	"github.com/jessevdk/go-flags"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"

	"go-code-challenge/apiserver"
)

type Args struct {
	Address  string `short:"a" long:"address" description:"The address to listen on for HTTP requests" default:"0.0.0.0"`
	Port     int    `short:"p" long:"port" description:"The port to listen on for HTTP requests" default:"3333"`
	Routes   bool   `short:"r" long:"routes" description:"Generate router documentation"`
	Database bool   `short:"d" long:"database" description:"Use a database"`
}

type Starship struct {
	args Args
	mode string
}

func NewStarship() *Starship {
	return &Starship{}
}

type StarshipBuilder interface {
	setConfig()
	setDatabase()
	setRepositories()
	setServices()
	setWebServer()
}

type BuildDirector struct {
	builder StarshipBuilder
}

func NewStarshipBuilder(sb StarshipBuilder) *BuildDirector {
	return &BuildDirector{
		builder: sb,
	}
}

func (sbd *BuildDirector) BuildStarship() {
	sbd.builder.setConfig()
	sbd.builder.setDatabase()
	sbd.builder.setRepositories()
	sbd.builder.setServices()
	sbd.builder.setWebServer()
}

func (star *Starship) setConfig() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	parser := flags.NewParser(&star.args, flags.Default)
	_, err := parser.Parse()
	if err != nil {
		log.Err(err)
		return
	}
}

func (star *Starship) setWebServer() {
	log.Info().Msg("Starting Web Server...")

	port := fmt.Sprintf(":%d", star.args.Port)
	webServer := apiserver.NewServer()

	router := chi.NewRouter()
	webServer.SetupRoutes(star.mode, router)

	if star.args.Routes {
		log.Info().Msg(docgen.JSONRoutesDoc(router))

		return
	}

	if err := http.ListenAndServe(port, router); err != nil {
		panic(err)
	}
}

func (star *Starship) setDatabase() {
}

func (star *Starship) setRepositories() {
}

func (star *Starship) setServices() {
}
