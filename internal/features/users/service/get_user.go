package users_service

import (
	"context"
	"fmt"

	"github.com/MrLaplace19/ToDoList/internal/core/domain"
)

func (s *UsersService) GetUser(
	ctx context.Context,
	userId int,
) (domain.User, error) {
	user, err := s.usersRepository.GetUser(ctx, userId)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user from repository: %w", err)
	}

	return user, nil
}
