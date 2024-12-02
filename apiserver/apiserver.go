package apiserver

import (
	"fmt"
	"github.com/go-chi/render"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
	httpSwagger "github.com/swaggo/http-swagger/v2"

	"go-code-challenge/apiserver/handlers"
	"go-code-challenge/internal"
)

type ApiServer struct {
	actionService internal.ActionServiceInterface
	usersService  internal.UserServiceInterface
}

func NewServer(actionService internal.ActionServiceInterface, usersService internal.UserServiceInterface) *ApiServer {
	return &ApiServer{
		actionService: actionService,
		usersService:  usersService,
	}
}

func (as *ApiServer) SetupRoutes(router *chi.Mux) {
	as.setupMiddleware(router)

	healthHandler := handlers.NewHealthHandler()
	actionsHandler := handlers.NewActionHandler(as.actionService)
	usershandler := handlers.NewUserHandler(as.usersService)

	as.registerCommonAPI(router, healthHandler)
	as.registerUsers(router, usershandler)
	as.registerActions(router, actionsHandler)

	serveSwagger(router)
}

func (as *ApiServer) registerCommonAPI(subrouter chi.Router, healthHandler *handlers.HealthHandler) {
	subrouter.Group(func(r chi.Router) {
		r.Get("/health", healthHandler.CheckHealth)
	})
}

func (as *ApiServer) registerUsers(subrouter chi.Router, userHandler *handlers.UserHandler) {
	subrouter.Group(func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Get("/{id}", userHandler.GetUser)
		})
	})
}

func (as *ApiServer) registerActions(subrouter chi.Router, actionsHandler *handlers.ActionHandler) {
	subrouter.Group(func(r chi.Router) {
		r.Route("/actions", func(r chi.Router) {
			r.Get("/{userID}/count", actionsHandler.GetActionCount)
			r.Get("/{type}/next", actionsHandler.GetNextActionProbabilities)
			r.Get("/referrals", actionsHandler.GetReferralIndex)
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

func (a *ApiServer) setupMiddleware(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))
}
