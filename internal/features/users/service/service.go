package users_service

import (
	"context"

	"github.com/Mommsent/todoapp-Studying.git/internal/core/domain"
)

type UserService struct {
	usersRepository UsersRepository
}

type UsersRepository interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUsers(ctx context.Context, limit, offset *int) ([]domain.User, error)
	GetUser(ctx context.Context, id int) (domain.User, error)
	DeleteUser(ctx context.Context, id int) error
	PatchUser(ctx context.Context, id int, user domain.User) (domain.User, error)
}

func NewUsersService(userRepository UsersRepository) *UserService {
	return &UserService{
		usersRepository: userRepository,
	}
}
