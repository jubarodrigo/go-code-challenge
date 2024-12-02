package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"

	"go-code-challenge/internal"
)

type UserHandler struct {
	userService internal.UserServiceInterface
}

func NewUserHandler(userService internal.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// GetUser @Summary Get user by ID
// @Description Retrieves a user's information given their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} users.User
// @Failure 400 {object} ErrResponse "Invalid user ID format"
// @Failure 404 {object} ErrResponse "User not found"
// @Router /users/{id} [get]
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Error().Msgf("Error converting user id to int: %v", err)
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	user, err := h.userService.FindUserByID(userID)
	if err != nil {
		log.Error().Msgf("Error getting user: %v", err)
		render.Render(w, r, ErrNotFound)
		return
	}

	render.JSON(w, r, user)
}
