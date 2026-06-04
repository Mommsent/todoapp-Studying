package users_service

import (
	"context"
	"fmt"

	"github.com/Mommsent/todoapp-Studying.git/internal/core/domain"
)

func (userService *UserService) PatchUser(ctx context.Context, id int, patch domain.UserPatch) (domain.User, error) {
	user, err := userService.usersRepository.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user: %w", err)
	}

	if err := user.ApplyPatch(patch); err != nil {
		return domain.User{}, fmt.Errorf("apply user patch: %w", err)
	}

	patchedUser, err := userService.usersRepository.PatchUser(ctx, id, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("patch user: %w", err)
	}

	return patchedUser, nil
}
