//go:generate mockgen -source=./interfaces.go -package=mocks -destination=../test/mocks/mock_services.go
package internal

import "go-code-challenge/internal/users"

type UserServiceInterface interface {
	FindUserByID(id int) (*users.User, error)
}

type ActionServiceInterface interface {
	FindActionCountByUserID(userID int) (int, error)
	FindNextActionProbabilities(actionType string) (map[string]float64, error)
	FindReferralIndex() (map[int]int, error)
}
