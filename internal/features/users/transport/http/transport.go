package users_transport_http

import (
	"context"
	"net/http"

	"github.com/Mommsent/todoapp-Studying.git/internal/core/domain"
	core_http_server "github.com/Mommsent/todoapp-Studying.git/internal/core/transport/http/server"
)

type UsersHTTPHandler struct {
	userService UserService
}

type UserService interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUsers(ctx context.Context, limit, offset *int) ([]domain.User, error)
	GetUser(ctx context.Context, id int) (domain.User, error)
	DeleteUser(ctx context.Context, id int) error
	PatchUser(ctx context.Context, id int, patch domain.UserPatch) (domain.User, error)
}

func NewUserHTTPHandler(userService UserService) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		userService: userService,
	}
}

func (handler *UsersHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: handler.CreateUser,
		},
		{
			Method:  http.MethodGet,
			Path:    "/users",
			Handler: handler.GetUsers,
		},
		{
			Method:  http.MethodGet,
			Path:    "/users/{id}",
			Handler: handler.GetUser,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/users/{id}",
			Handler: handler.DeleteUser,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/users/{id}",
			Handler: handler.PatchUser,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/users/{id}",
			Handler: handler.PatchUser,
		},
	}
}
