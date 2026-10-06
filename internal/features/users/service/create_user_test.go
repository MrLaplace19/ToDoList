package users_service

import (
	"context"
	"errors"
	"testing"

	"github.com/MrLaplace19/ToDoList/internal/core/domain"
	core_errors "github.com/MrLaplace19/ToDoList/internal/core/errors"
)

type fakeUserRepository struct {
	UsersRepository
	createCalled bool
}

func (f *fakeUserRepository) CreateUser(
	ctx context.Context,
	user domain.User,
) (domain.User, error) {
	f.createCalled = true
	return user, nil
}

func TestUserServiceCreate_InvalidUser(t *testing.T) {
	repository := &fakeUserRepository{}
	service := NewUserService(repository)
	user := domain.NewUserUninitialized("ab", nil)
	_, err := service.CreateUser(context.Background(), user)

	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, core_errors.ErrInvalidArgument) {
		t.Fatal("expected ErrInvalidArgument: %w", err)
	}
	if repository.createCalled {
		t.Fatal("repository should be called invalid user")
	}
}
