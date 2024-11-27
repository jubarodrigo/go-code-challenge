package handlers

import (
	"net/http"

	"github.com/go-chi/render"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

type HealthStatus struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type healthStatusResponse struct {
	*HealthStatus
}

func (ur *HealthStatus) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func newDeviceResponse(device *HealthStatus) *healthStatusResponse {
	return &healthStatusResponse{HealthStatus: device}
}

func (h *HealthHandler) CheckHealth(w http.ResponseWriter, r *http.Request) {
	response := &HealthStatus{
		Status:  "healthy",
		Message: "server up and running",
	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, newDeviceResponse(response))
}
