package services

import (
	"fmt"

	"go-code-challenge/datastore"
	"go-code-challenge/internal"
)

const referUser = "REFER_USER"

type ActionService struct {
	repo datastore.DatasJsonRepositoryInterface
}

func NewActionService(repo datastore.DatasJsonRepositoryInterface) internal.ActionServiceInterface {
	return &ActionService{repo: repo}
}

func (s *ActionService) FindActionCountByUserID(userID int) (int, error) {
	actions, err := s.repo.GetActionsByUserID(userID)
	if err != nil {
		return 0, fmt.Errorf("error getting actions: %w", err)
	}
	return len(actions), nil
}

func (s *ActionService) FindNextActionProbabilities(actionType string) (map[string]float64, error) {
	actions, err := s.repo.GetAllActions()
	if err != nil {
		return nil, fmt.Errorf("error getting actions: %w", err)
	}

	nextActionCounts := make(map[string]int)
	totalCount := 0

	for i := 0; i < len(actions)-1; i++ {
		if actions[i].Type == actionType {
			nextActionType := actions[i+1].Type
			nextActionCounts[nextActionType]++
			totalCount++
		}
	}

	probabilities := make(map[string]float64)
	if totalCount > 0 {
		for action, count := range nextActionCounts {
			probabilities[action] = float64(count) / float64(totalCount)
		}
	}

	return probabilities, nil
}

func (s *ActionService) FindReferralIndex() (map[int]int, error) {
	actions, err := s.repo.GetAllActions()
	if err != nil {
		return nil, fmt.Errorf("error getting actions: %w", err)
	}

	referralGraph := make(map[int][]int)
	for _, action := range actions {
		if action.Type == referUser {
			referralGraph[action.UserID] = append(referralGraph[action.UserID], action.TargetUser)
		}
	}

	referralIndex := make(map[int]int)

	var calculateReferrals func(int, map[int]bool) int
	calculateReferrals = func(userID int, visited map[int]bool) int {
		if visited[userID] {
			return 0
		}
		visited[userID] = true

		count := 0
		for _, referredUser := range referralGraph[userID] {
			count++                                            // Direct referral
			count += calculateReferrals(referredUser, visited) // Indirect referrals
		}
		return count
	}

	for userID := range referralGraph {
		visited := make(map[int]bool)
		referralIndex[userID] = calculateReferrals(userID, visited)
	}

	return referralIndex, nil
}
