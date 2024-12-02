package services

import (
	"go-code-challenge/datastore"
	"go-code-challenge/internal/users"
)

type UserService struct {
	repo datastore.DatasJsonRepositoryInterface
}

func NewUserService(repo datastore.DatasJsonRepositoryInterface) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserByID(id int) (*users.User, error) {
	return s.repo.GetUserByID(id)
}
