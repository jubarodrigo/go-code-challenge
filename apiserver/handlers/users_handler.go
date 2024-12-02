package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

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

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	user, err := h.userService.FindUserByID(userID)
	if err != nil {
		render.Render(w, r, ErrNotFound)
		return
	}

	render.JSON(w, r, user)
}
