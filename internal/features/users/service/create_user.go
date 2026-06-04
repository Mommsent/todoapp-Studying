package users_service

import (
	"context"
	"fmt"

	"github.com/Mommsent/todoapp-Studying.git/internal/core/domain"
)

func (userService *UserService) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	if err := user.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("Validate user domain: %w", err)
	}

	user, err := userService.usersRepository.CreateUser(ctx, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("create user: %w", err)
	}

	return user, nil
}
