package users_service

import (
	"context"
	"fmt"

	"github.com/Mommsent/todoapp-Studying.git/internal/core/domain"
)

func (userService *UserService) GetUser(ctx context.Context, id int) (domain.User, error) {
	user, err := userService.usersRepository.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user from repository: %w", err)
	}

	return user, nil
}
