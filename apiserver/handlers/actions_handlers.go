package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"go-code-challenge/internal"
)

type ActionHandler struct {
	actionService internal.ActionServiceInterface
}

func NewActionHandler(actionService internal.ActionServiceInterface) *ActionHandler {
	return &ActionHandler{
		actionService: actionService,
	}
}

func (h *ActionHandler) GetActionCount(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "userID"))
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	count, err := h.actionService.FindActionCountByUserID(userID)
	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}

	render.JSON(w, r, map[string]int{"count": count})
}

func (h *ActionHandler) GetNextActionProbabilities(w http.ResponseWriter, r *http.Request) {
	actionType := chi.URLParam(r, "type")
	if actionType == "" {
		render.Render(w, r, ErrInvalidRequest(fmt.Errorf("action type is required")))
		return
	}

	probabilities, err := h.actionService.FindNextActionProbabilities(actionType)
	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}

	render.JSON(w, r, probabilities)
}

func (h *ActionHandler) GetReferralIndex(w http.ResponseWriter, r *http.Request) {
	referralIndex, err := h.actionService.FindReferralIndex()
	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}

	render.JSON(w, r, referralIndex)
}
