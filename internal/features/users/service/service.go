package users_service

import (
	"context"

	"github.com/MrLaplace19/ToDoList/internal/core/domain"
)


type UsersService struct{
	usersRepository UsersRepository
}

type UsersRepository interface{
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

func NewUserService(usersRepository UsersRepository)*UsersService{
	return &UsersService{
		usersRepository: usersRepository,
	}
}

