package service

import (
	"context"
	"errors"
	"user-service/internal/app/domain"
)

//go:generate moq -out service_user_mock_test.go . RepositoryUser UtilBcrypt UtilJwt
type RepositoryUser interface {
	GetByID(ctx context.Context, id int64) domain.Repository
	GetByEmail(ctx context.Context, email string) domain.Repository
	Create(ctx context.Context, user *domain.User) domain.Repository
}

type UtilBcrypt interface {
	CheckIfPasswordHashIsEqual(repoPassword []byte, paramPassword []byte) error
	GenerateHashFromPlainPassword(paramPassword []byte) (string, error)
}

type UtilJwt interface {
	GenerateJWT(userID int64) (string, error)
}

type ServiceUser struct {
	repositoryUser RepositoryUser
	utilBcrypt     UtilBcrypt
	utilJwt        UtilJwt
}

func NewServiceUser(repositoryUser RepositoryUser, utilBcrypt UtilBcrypt, utilJwt UtilJwt) *ServiceUser {
	return &ServiceUser{repositoryUser, utilBcrypt, utilJwt}
}

func (su ServiceUser) Login(ctx context.Context, param domain.LoginParam) domain.Service {
	repo := su.repositoryUser.GetByEmail(ctx, param.Email)
	if repo.Error != nil {
		return domain.Service{
			Hash:  "GM5EQNy8msHL",
			Error: errors.New("email not exist"),
		}
	}

	err := su.utilBcrypt.CheckIfPasswordHashIsEqual([]byte(repo.User.Password), []byte(param.Password))
	if err != nil {
		return domain.Service{
			Hash:  "GMz3Y5sNe96O",
			Error: errors.New("wrong email or password"),
		}
	}

	token, err := su.utilJwt.GenerateJWT(repo.User.ID)
	if err != nil {
		return domain.Service{
			Hash:  "GMItRDipgI19",
			Error: err,
		}
	}

	return domain.Service{
		Hash:        "GMijx79VQ7bS",
		RawResponse: token,
	}
}

func (su ServiceUser) Register(ctx context.Context, param domain.RegisterParam) domain.Service {
	hashedPassword, err := su.utilBcrypt.GenerateHashFromPlainPassword([]byte(param.Password))
	if err != nil {
		return domain.Service{
			Hash:  "GMeLrJKgIHO9",
			Error: errors.New("failed generate hash password"),
		}
	}

	user := domain.User{
		Name:     param.Name,
		Email:    param.Email,
		Password: hashedPassword,
	}
	repo := su.repositoryUser.Create(ctx, &user)
	if repo.Error != nil {
		return domain.Service{
			Hash:  "GMHWzXXMSk9p",
			Error: errors.New("failed create user"),
		}
	}

	token, err := su.utilJwt.GenerateJWT(user.ID)
	if err != nil {
		return domain.Service{
			Hash:  "GMlpLJtRdVdx",
			Error: err,
		}
	}

	return domain.Service{
		Hash:        "GMmmtY5oiY9J",
		RawResponse: token,
	}
}

func (su ServiceUser) GetByID(ctx context.Context, id int64) domain.Service {
	if id <= 0 {
		return domain.Service{
			Hash:  "GML8UxsMd5E5",
			Error: errors.New("invalid user id"),
		}
	}

	repo := su.repositoryUser.GetByID(ctx, id)
	if repo.Error != nil {
		return domain.Service{
			Hash:  "GMO9escS9esE",
			Error: errors.New("user not found"),
		}
	}

	return domain.Service{
		Hash: "GMDxylSw7Gnm",
		User: repo.User,
	}
}
