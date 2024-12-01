package apiserver

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"go-code-challenge/apiserver/handlers"
	"net/http"
)

type ApiServer struct {
}

func NewServer() *ApiServer {
	return &ApiServer{}
}

func (a *ApiServer) SetupRoutes(router *chi.Mux) {
	healthHandler := handlers.NewHealthHandler()

	a.registerCommonAPI(router, healthHandler)
	a.registerUsers(router)
	a.registerActions(router)

	serveSwagger(router)
}

func (a *ApiServer) registerCommonAPI(subrouter chi.Router, healthHandler *handlers.HealthHandler) {
	subrouter.Group(func(r chi.Router) {
		r.Get("/health", healthHandler.CheckHealth)
	})
}

func (a *ApiServer) registerUsers(subrouter chi.Router, healthHandler *handlers.HealthHandler) {
	subrouter.Group(func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Get("/{id}", healthHandler.CheckHealth)
		})
	})
}

func (a *ApiServer) registerActions(subrouter chi.Router, healthHandler *handlers.HealthHandler) {
	subrouter.Group(func(r chi.Router) {
		r.Route("/actions", func(r chi.Router) {
			r.Get("/{userID}/count", healthHandler.CheckHealth)
			r.Get("/{type}/next", healthHandler.CheckHealth)
			r.Get("/{userID}", healthHandler.CheckHealth)
			r.Get("/referrals", healthHandler.CheckHealth)
		})
	})
}

func getSpecs() map[string]string {
	return map[string]string{
		"App": "docs/swagger.yaml",
	}
}

func serveSwagger(router chi.Router) {
	log.Info().Msg("serving swagger at /swagger/index.html")

	// serve the docs folder at /swagger/docs/
	fs := http.FileServer(http.Dir("docs"))
	router.Handle("/swagger/docs/*", http.StripPrefix("/swagger/docs/", fs))

	// server swagger ui with each swagger spec in getSpecs() configured
	specUrls := ""
	for specName, specUrl := range getSpecs() {
		specUrls += fmt.Sprintf(`{name: "%s", url: "%s"},`, specName, specUrl)
	}
	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.UIConfig(map[string]string{
			"urls": fmt.Sprintf("[%s]", specUrls),
		}),
	))
}
