package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/MrLaplace19/ToDoList/internal/core/domain"
	core_errors "github.com/MrLaplace19/ToDoList/internal/core/errors"
	"github.com/jackc/pgx/v5"
)

func (r *UsersRepository) GetUser(
		ctx context.Context,
		userId int,
	)(domain.User, error){
		ctx, cancel := context.WithTimeout(ctx, r.pool.OptionalTimeOut())
		defer cancel()

		query := `
		SELECT id,version,full_name,phone_number
		FROM todoapp.users
		WHERE id=$1
		`

		row := r.pool.QueryRow(ctx, query, userId)

		var userModel UserModel
		err := row.Scan(
			&userModel.ID,
			&userModel.Version,
			&userModel.FullName,
			&userModel.PhoneNumber,
		)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows){
				return domain.User{}, fmt.Errorf("user with id='%d': %w", userId, core_errors.ErrNotFound)
			}
			return domain.User{}, fmt.Errorf("scan error: %w", err)
		}

		userDomain := domain.User{
			ID: userModel.ID,
			Version: userModel.Version,
			FullName: userModel.FullName,
			PhoneNumber: userModel.PhoneNumber,
		}

		return userDomain, nil
	}