package datastore

import (
	"go-code-challenge/internal/actions"
	"go-code-challenge/internal/users"
)

type UserRepository interface {
	FindByID(id int) (users.User, error)
}

type ActionRepository interface {
	FindByUserID(userID int) ([]actions.Action, error)
	FindAll() ([]actions.Action, error)
}
