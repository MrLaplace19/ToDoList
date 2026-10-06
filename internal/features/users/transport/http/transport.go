package users_transport_http

import (
	"context"
	"net/http"

	"github.com/MrLaplace19/ToDoList/internal/core/domain"
	core_http_server "github.com/MrLaplace19/ToDoList/internal/core/transport/http/server"
)

type UserHttpHandler struct {
	userService UserService
}

type UserService interface {
	CreateUser(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)
	GetUsers(
		ctx context.Context,
		limit *int,
		offset *int,
	) ([]domain.User, error)
	GetUser(
		ctx context.Context,
		userId int,
	) (domain.User, error)
	DeleteUser(
		ctx context.Context,
		userID int,
	) error
	PatchUser(
		ctx context.Context,
		userID int,
		userPatch domain.UserPatch,
	) (domain.User, error)
}

func NewUserHttpHandler(userService UserService) *UserHttpHandler {
	return &UserHttpHandler{
		userService: userService,
	}
}

func (h *UserHttpHandler) Routes() []core_http_server.Router {
	return []core_http_server.Router{
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: h.CreateUser,
		},
		{
			Method:  http.MethodGet,
			Path:    "/users",
			Handler: h.GetUsers,
		},
		{
			Method:  http.MethodGet,
			Path:    "/users/{id}",
			Handler: h.GetUser,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/users/{id}",
			Handler: h.DeleteUser,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/users/{id}",
			Handler: h.PatchUser,
		},
	}
}
