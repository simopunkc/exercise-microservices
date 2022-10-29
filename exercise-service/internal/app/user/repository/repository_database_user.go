package repository

import (
	"context"
	"exercise-service/internal/app/domain"

	"gorm.io/gorm"
)

type RepositoryDatabaseUser struct {
	db *gorm.DB
}

func NewRepositoryDatabaseUser(db *gorm.DB) *RepositoryDatabaseUser {
	return &RepositoryDatabaseUser{db}
}

func (rdu RepositoryDatabaseUser) IsUserExists(ctx context.Context, userID int64) bool {
	var user domain.User
	err := rdu.db.WithContext(ctx).Where("id = ?", userID).Take(&user).Error
	return err != nil && user.ID > 0
}
