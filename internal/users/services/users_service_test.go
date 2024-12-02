package services

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"go-code-challenge/internal/users"
	"go-code-challenge/test/mocks"
)

func TestFindUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDatasJsonRepositoryInterface(ctrl)
	service := NewUserService(mockRepo)

	testTime := time.Now()

	tests := []struct {
		name        string
		id          int
		mockUser    *users.User
		mockErr     error
		expectedErr error
	}{
		{
			name: "successful user retrieval",
			id:   1,
			mockUser: &users.User{
				ID:        1,
				Name:      "John Doe",
				CreatedAt: testTime,
			},
			mockErr:     nil,
			expectedErr: nil,
		},
		{
			name:        "user not found",
			id:          2,
			mockUser:    nil,
			mockErr:     errors.New("user not found"),
			expectedErr: errors.New("user not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().
				GetUserByID(tt.id).
				Return(tt.mockUser, tt.mockErr)

			user, err := service.FindUserByID(tt.id)

			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErr.Error(), err.Error())
				assert.Nil(t, user)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.mockUser, user)
			}
		})
	}
}
