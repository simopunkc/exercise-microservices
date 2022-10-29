package service

import (
	"context"
	"time"
)

//go:generate moq -out service_user_mock_test.go . RepositoryUser
type RepositoryUser interface {
	IsUserExists(ctx context.Context, userID int64) bool
}

type ServiceUser struct {
	repositoryUser RepositoryUser
}

func NewServiceUser(repositoryUser RepositoryUser) *ServiceUser {
	return &ServiceUser{
		repositoryUser: repositoryUser,
	}
}

func (su ServiceUser) IsUserExists(ctx context.Context, userID int64) bool {
	if userID == 0 {
		return false
	}
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return su.repositoryUser.IsUserExists(ctx, userID)
}
