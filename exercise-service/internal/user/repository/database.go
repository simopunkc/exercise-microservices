package repository

import (
	"context"
	"exercise-service/internal/domain"

	"gorm.io/gorm"
)

type DatabaseRepo struct {
	db *gorm.DB
}

func NewDatabaseRepo(db *gorm.DB) *DatabaseRepo {
	return &DatabaseRepo{db}
}

func (dr DatabaseRepo) IsUserExists(ctx context.Context, userID int) bool {
	var user domain.User
	err := dr.db.WithContext(ctx).Where("id = ?", userID).Take(&user).Error
	return err != nil && user.ID > 0
}
