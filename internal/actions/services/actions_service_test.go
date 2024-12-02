package services

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"go-code-challenge/internal/actions"
	"go-code-challenge/test/mocks"
)

func TestFindActionCountByUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDatasJsonRepositoryInterface(ctrl)
	service := NewActionService(mockRepo)

	testTime := time.Now()

	tests := []struct {
		name          string
		userID        int
		mockActions   []actions.Action
		mockErr       error
		expectedCount int
		expectedErr   error
	}{
		{
			name:   "user with actions",
			userID: 1,
			mockActions: []actions.Action{
				{ID: 1, UserID: 1, Type: "ADD_TO_CRM", CreatedAt: testTime},
				{ID: 2, UserID: 1, Type: "REFER_USER", CreatedAt: testTime},
			},
			mockErr:       nil,
			expectedCount: 2,
			expectedErr:   nil,
		},
		{
			name:          "user without actions",
			userID:        2,
			mockActions:   []actions.Action{},
			mockErr:       nil,
			expectedCount: 0,
			expectedErr:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().
				GetActionsByUserID(tt.userID).
				Return(tt.mockActions, tt.mockErr)

			count, err := service.FindActionCountByUserID(tt.userID)

			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErr.Error(), err.Error())
				assert.Equal(t, tt.expectedCount, count)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedCount, count)
			}
		})
	}
}

func TestFindNextActionProbabilities(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDatasJsonRepositoryInterface(ctrl)
	service := NewActionService(mockRepo)

	testTime := time.Now()

	tests := []struct {
		name          string
		actionType    string
		mockActions   []actions.Action
		mockErr       error
		expectedProbs map[string]float64
		expectedErr   error
	}{
		{
			name:       "valid action sequence",
			actionType: "ADD_TO_CRM",
			mockActions: []actions.Action{
				{ID: 1, Type: "ADD_TO_CRM", CreatedAt: testTime},
				{ID: 2, Type: "REFER_USER", CreatedAt: testTime},
				{ID: 3, Type: "ADD_TO_CRM", CreatedAt: testTime},
				{ID: 4, Type: "VIEW_CONVERSATION", CreatedAt: testTime},
			},
			mockErr: nil,
			expectedProbs: map[string]float64{
				"REFER_USER":        0.5,
				"VIEW_CONVERSATION": 0.5,
			},
			expectedErr: nil,
		},
		{
			name:          "no actions found",
			actionType:    "NONEXISTENT",
			mockActions:   []actions.Action{},
			mockErr:       nil,
			expectedProbs: map[string]float64{},
			expectedErr:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().
				GetAllActions().
				Return(tt.mockActions, tt.mockErr)

			probs, err := service.FindNextActionProbabilities(tt.actionType)

			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErr.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedProbs, probs)
			}
		})
	}
}

func TestFindReferralIndex(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDatasJsonRepositoryInterface(ctrl)
	service := NewActionService(mockRepo)

	testTime := time.Now()

	tests := []struct {
		name          string
		mockActions   []actions.Action
		mockErr       error
		expectedIndex map[int]int
		expectedErr   error
	}{
		{
			name: "valid referral chain",
			mockActions: []actions.Action{
				{ID: 1, Type: "REFER_USER", UserID: 1, TargetUser: 2, CreatedAt: testTime},
				{ID: 2, Type: "REFER_USER", UserID: 2, TargetUser: 3, CreatedAt: testTime},
				{ID: 3, Type: "REFER_USER", UserID: 1, TargetUser: 4, CreatedAt: testTime},
			},
			mockErr: nil,
			expectedIndex: map[int]int{
				1: 3, // Referred users 2 and 4, and user 2 referred 3
				2: 1, // Referred user 3
			},
			expectedErr: nil,
		},
		{
			name: "cyclic referral",
			mockActions: []actions.Action{
				{ID: 1, Type: "REFER_USER", UserID: 1, TargetUser: 2, CreatedAt: testTime},
				{ID: 2, Type: "REFER_USER", UserID: 2, TargetUser: 1, CreatedAt: testTime},
			},
			mockErr: nil,
			expectedIndex: map[int]int{
				1: 2, // Counts both direct reference to 2 and indirect reference to 1 through 2
				2: 2, // Counts both direct reference to 1 and indirect reference to 2 through 1
			},
			expectedErr: nil,
		},
		{
			name: "no referrals",
			mockActions: []actions.Action{
				{ID: 1, Type: "ADD_TO_CRM", UserID: 1, CreatedAt: testTime},
				{ID: 2, Type: "VIEW_CONVERSATION", UserID: 2, CreatedAt: testTime},
			},
			mockErr:       nil,
			expectedIndex: map[int]int{},
			expectedErr:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().
				GetAllActions().
				Return(tt.mockActions, tt.mockErr)

			index, err := service.FindReferralIndex()

			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErr.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedIndex, index)
			}
		})
	}
}
