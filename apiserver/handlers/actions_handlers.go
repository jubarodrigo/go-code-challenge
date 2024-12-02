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

// GetActionCount @Summary Get action count by user ID
// @Description Returns the total number of actions performed by a specific user
// @Tags actions
// @Accept json
// @Produce json
// @Param userID path int true "User ID"
// @Success 200 {object} map[string]int "count"
// @Failure 400 {object} ErrResponse "Invalid user ID format"
// @Failure 500 {object} ErrResponse "Internal server error"
// @Router /actions/{userID}/count [get]
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

// GetNextActionProbabilities @Summary Get next action probabilities
// @Description Returns probability distribution of next actions given a specific action type
// @Tags actions
// @Accept json
// @Produce json
// @Param type path string true "Action Type"
// @Success 200 {object} map[string]float64
// @Failure 400 {object} ErrResponse "Invalid or missing action type"
// @Failure 500 {object} ErrResponse "Internal server error"
// @Router /actions/{type}/next [get]
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

// GetReferralIndex @Summary Get referral index
// @Description Returns the referral index for all users, showing direct and indirect referrals
// @Tags actions
// @Accept json
// @Produce json
// @Success 200 {object} map[int]int
// @Failure 500 {object} ErrResponse "Internal server error"
// @Router /actions/referrals [get]
func (ah *ActionHandler) GetReferralIndex(w http.ResponseWriter, r *http.Request) {
	referralIndex, err := ah.actionService.FindReferralIndex()
	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}

	render.JSON(w, r, referralIndex)
}
