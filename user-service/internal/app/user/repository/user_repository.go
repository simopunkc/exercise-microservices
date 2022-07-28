package repository

import (
	"context"
	"user-service/internal/app/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur UserRepository) GetByID(ctx context.Context, id int) (domain.User, error) {
	var user domain.User
	err := ur.db.WithContext(ctx).Where("id = ?", id).Take(&user).Error
	return user, err
}

func (ur UserRepository) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	var user domain.User
	err := ur.db.WithContext(ctx).Where("email = ?", email).Take(&user).Error
	return user, err
}

func (ur UserRepository) Create(ctx context.Context, user *domain.User) error {
	return ur.db.WithContext(ctx).Create(user).Error
}
