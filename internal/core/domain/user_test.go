package domain

import (
	"errors"
	"strings"
	"testing"

	core_errors "github.com/MrLaplace19/ToDoList/internal/core/errors"
)

func TestUserValidate(t *testing.T) {
	validatePhone := "+79998887766"
	invalidatePhone := "73474378"

	tests := []struct {
		name    string
		user    User
		wantErr bool
	}{
		{
			name:    "valid user",
			user:    NewUserUninitialized("IVAN", &validatePhone),
			wantErr: false,
		},
		{
			name:    "invalid phone",
			user:    NewUserUninitialized("ANTON", &invalidatePhone),
			wantErr: true,
		},
		{
			name:    "invalid name",
			user:    NewUserUninitialized("ad", nil),
			wantErr: true,
		},
		{
			name:    "full name with minimal symbols",
			user:    NewUserUninitialized("ABC", &validatePhone),
			wantErr: false,
		},
		{
			name:    "full name with maximum symbols failed",
			user:    NewUserUninitialized(strings.Repeat("a", 101), &validatePhone),
			wantErr: true,
		},
		{
			name:    "phone is optional",
			user:    NewUserUninitialized("STEPAN", nil),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.user.Validate()
			if tt.wantErr && err == nil {
				t.Fatal("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Fatal("expected no error, got %w", err)
			}
			if tt.wantErr && !errors.Is(err, core_errors.ErrInvalidArgument) {
				t.Fatal("expected ErrInvalidArgument, got %w", err)
			}
		})
	}
}
