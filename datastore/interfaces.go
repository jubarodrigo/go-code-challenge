//go:generate mockgen -source=./interfaces.go -package=mocks -destination=../test/mocks/mock_repos.go
package datastore

import (
	"go-code-challenge/internal/actions"
	"go-code-challenge/internal/users"
)

type UserRepository interface {
	GetUserByID(id int) (*users.User, error)
}

type ActionRepository interface {
	GetActionsByUserID(userID int) ([]actions.Action, error)
	GetAllActions() ([]actions.Action, error)
}
