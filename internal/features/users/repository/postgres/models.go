package users_postgres_repository

import "github.com/Mommsent/todoapp-Studying.git/internal/core/domain"

type UserModel struct {
	ID          int
	Version     int
	FullName    string
	PhoneNumber *string
}

func usersDomainFromModels(users []UserModel) []domain.User {
	userDomains := make([]domain.User, len(users))

	for i, user := range users {
		userDomains[i] = domain.NewUser(
			user.ID,
			user.Version,
			user.FullName,
			user.PhoneNumber,
		)
	}

	return userDomains
}
