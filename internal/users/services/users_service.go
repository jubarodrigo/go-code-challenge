package services

import (
	"go-code-challenge/datastore"
	"go-code-challenge/internal"
	"go-code-challenge/internal/users"
)

type UserService struct {
	repo datastore.DatasJsonRepositoryInterface
}

func NewUserService(repo datastore.DatasJsonRepositoryInterface) internal.UserServiceInterface {
	return &UserService{repo: repo}
}

func (s *UserService) FindUserByID(id int) (*users.User, error) {
	return s.repo.GetUserByID(id)
}
