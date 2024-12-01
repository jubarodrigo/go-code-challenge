//go:generate mockgen -source=./interfaces.go -package=mocks -destination=../test/mocks/mock_services.go
package internal

import "go-code-challenge/internal/users"

type UserService interface {
	FindUserByID(id int) (users.User, error)
}

type ActionService interface {
	FindActionCountByUserID(userID int) (int, error)
	FindNextActionProbabilities(actionType string) (map[string]float64, error)
	FindReferralIndex() (map[int]int, error)
}
