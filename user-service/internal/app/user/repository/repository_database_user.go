package repository

import (
	"context"
	"user-service/internal/app/domain"

	"gorm.io/gorm"
)

type RepositoryUser struct {
	db *gorm.DB
}

func NewRepositoryDatabaseUser(db *gorm.DB) *RepositoryUser {
	return &RepositoryUser{db}
}

func (ru RepositoryUser) GetByID(ctx context.Context, id int64) domain.Repository {
	var user domain.User
	err := ru.db.WithContext(ctx).Where("id = ?", id).Take(&user).Error
	if err != nil {
		return domain.Repository{
			Error: err,
		}
	}
	return domain.Repository{
		User: user,
	}
}

func (ru RepositoryUser) GetByEmail(ctx context.Context, email string) domain.Repository {
	var user domain.User
	err := ru.db.WithContext(ctx).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return domain.Repository{
			Error: err,
		}
	}
	return domain.Repository{
		User: user,
	}
}

func (ru RepositoryUser) Create(ctx context.Context, user *domain.User) domain.Repository {
	err := ru.db.WithContext(ctx).Create(user).Error
	return domain.Repository{
		Error: err,
	}
}
