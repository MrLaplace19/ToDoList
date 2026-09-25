package users_service

import (
	"context"
	"fmt"

	"github.com/MrLaplace19/ToDoList/internal/core/domain"
)

func (s *UsersService) PatchUser(
	ctx context.Context,
	userID int,
	userPatch domain.UserPatch,
) (domain.User, error) {
	user, err := s.usersRepository.GetUser(ctx, userID)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user: %w", err)
	}
	if err := user.ApplyPatch(userPatch); err != nil {
		return domain.User{}, fmt.Errorf("apply user patch: %w", err)
	}

	patchedUser, err := s.usersRepository.PatchUser(ctx, userID, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("path user to repository: %w", err)
	}

	return patchedUser, nil
}
