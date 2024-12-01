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

func (ah *ActionHandler) GetActionCount(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "userID"))
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	count, err := ah.actionService.FindActionCountByUserID(userID)
	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}

	render.JSON(w, r, map[string]int{"count": count})
}

func (ah *ActionHandler) GetNextActionProbabilities(w http.ResponseWriter, r *http.Request) {
	actionType := chi.URLParam(r, "type")
	if actionType == "" {
		render.Render(w, r, ErrInvalidRequest(fmt.Errorf("action type is required")))
		return
	}

	probabilities, err := ah.actionService.FindNextActionProbabilities(actionType)
	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}

	render.JSON(w, r, probabilities)
}

func (ah *ActionHandler) GetReferralIndex(w http.ResponseWriter, r *http.Request) {
	referralIndex, err := ah.actionService.FindReferralIndex()
	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}

	render.JSON(w, r, referralIndex)
}
