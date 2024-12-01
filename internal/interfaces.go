package internal

import "go-code-challenge/internal/users"

type UserService interface {
	GetUserByID(id int) (users.User, error)
}

type ActionService interface {
	GetActionCountByUserID(userID int) (int, error)
	GetNextActionProbabilities(actionType string) (map[string]float64, error)
	GetReferralIndex() (map[int]int, error)
}
