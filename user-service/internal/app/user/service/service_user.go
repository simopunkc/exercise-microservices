package service

import (
	"context"
	"errors"
	"user-service/internal/app/domain"
	"user-service/internal/pkg/util"

	"golang.org/x/crypto/bcrypt"
)

type RepositoryUser interface {
	GetByID(ctx context.Context, id int64) domain.Repository
	GetByEmail(ctx context.Context, email string) domain.Repository
	Create(ctx context.Context, user *domain.User) domain.Repository
}

type ServiceUser struct {
	repositoryUser RepositoryUser
}

func NewServiceUser(repositoryUser RepositoryUser) *ServiceUser {
	return &ServiceUser{repositoryUser}
}

func (su ServiceUser) Login(ctx context.Context, param domain.LoginParam) domain.Service {
	repo := su.repositoryUser.GetByEmail(ctx, param.Email)
	if repo.Error != nil {
		return domain.Service{
			Error: errors.New("wrong email or password"),
		}
	}

	err := bcrypt.CompareHashAndPassword([]byte(repo.User.Password), []byte(param.Password))
	if err != nil {
		return domain.Service{
			Error: errors.New("wrong email or password"),
		}
	}

	token, err := util.GenerateJWT(repo.User.ID)
	if err != nil {
		return domain.Service{
			Error: err,
		}
	}

	return domain.Service{
		RawResponse: token,
	}
}

func (su ServiceUser) Register(ctx context.Context, param domain.RegisterParam) domain.Service {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(param.Password), bcrypt.DefaultCost)
	user := domain.User{
		Name:     param.Name,
		Email:    param.Email,
		Password: string(hashedPassword),
	}
	repo := su.repositoryUser.Create(ctx, &user)
	if repo.Error != nil {
		return domain.Service{
			Error: errors.New("failed create user"),
		}
	}

	token, err := util.GenerateJWT(user.ID)
	if err != nil {
		return domain.Service{
			Error: err,
		}
	}

	return domain.Service{
		RawResponse: token,
	}
}

func (su ServiceUser) GetByID(ctx context.Context, id int64) domain.Service {
	if id <= 0 {
		return domain.Service{
			Error: errors.New("invalid user id"),
		}
	}

	repo := su.repositoryUser.GetByID(ctx, id)
	if repo.Error != nil {
		return domain.Service{
			Error: errors.New("user not found"),
		}
	}

	return domain.Service{
		User: repo.User,
	}
}
