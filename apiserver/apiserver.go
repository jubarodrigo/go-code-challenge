package apiserver

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type ApiServer struct {
}

func NewServer() *ApiServer {
	return &ApiServer{}
}

func (a *ApiServer) SetupRoutes(envBaseUrl string, r *chi.Mux) {
	envBaseUrl = fmt.Sprintf("/%s", envBaseUrl)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/index.html")
	})

	fs := http.FileServer(http.Dir("web"))
	r.Get("/web/*", http.StripPrefix("/web", fs).ServeHTTP)

	a.registerCommonAPI(envBaseUrl, r)

	serveSwagger(r)
}

func handleServerTime(response http.ResponseWriter, request *http.Request) {
	log.Info().Msg(fmt.Sprintf("Handling Server time %s", time.Now().String()))
	type serverTime struct {
		Time string `json:"time"`
	}
	now := time.Now()
	data := &serverTime{
		Time: now.Format(time.RFC3339),
	}
	render.JSON(response, request, data)
}

func (a *ApiServer) registerCommonAPI(envBaseUrl string, subrouter chi.Router) {
	subrouter.Group(func(r chi.Router) {
		r.Get(envBaseUrl+"/server-time", handleServerTime)
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
