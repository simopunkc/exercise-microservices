package service

import (
	"context"
	"errors"
	"user-service/internal/app/domain"
)

func (us UserService) GetByID(ctx context.Context, id int) (domain.User, error) {
	if id <= 0 {
		return domain.User{}, errors.New("invalid user id")
	}
	return us.userRepo.GetByID(ctx, id)
}
