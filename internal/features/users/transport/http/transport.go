package users_transport_http

import (
	"context"
	"net/http"

	"github.com/MrLaplace19/ToDoList/internal/core/domain"
	core_http_server "github.com/MrLaplace19/ToDoList/internal/core/transport/http/server"
)

type UserHttpHandler struct{
	userService UserService
}

type UserService interface{
	CreateUser(
		ctx context.Context,
		user domain.User,
	)(domain.User, error)
	GetUser(
		ctx context.Context,
		limit *int,
		offset *int,
	)([]domain.User,error)
}

func NewUserHttpHandler(userService UserService) *UserHttpHandler{
	return &UserHttpHandler{
		userService: userService,
	}
}

func (h *UserHttpHandler) Routes() []core_http_server.Router{
	return []core_http_server.Router{
		{
			Method: http.MethodPost,
			Path: "/users",
			Handler: h.CreateUser,
		},
		{
			Method: http.MethodGet,
			Path: "/users",
			Handler: h.GetUsers,
		},
	}
}
