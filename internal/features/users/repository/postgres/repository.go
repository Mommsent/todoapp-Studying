package users_postgres_repository

import core_posgres_pool "github.com/Mommsent/todoapp-Studying.git/internal/core/repository/postgres/pool"

type UsersRepository struct {
	pool core_posgres_pool.Pool
}

func NewUsersRepository(pool core_posgres_pool.Pool) *UsersRepository {
	return &UsersRepository{
		pool: pool,
	}
}
