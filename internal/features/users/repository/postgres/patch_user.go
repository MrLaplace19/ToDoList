package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/MrLaplace19/ToDoList/internal/core/domain"
	core_errors "github.com/MrLaplace19/ToDoList/internal/core/errors"
	"github.com/jackc/pgx/v5"
)

func (r *UsersRepository) PatchUser(
	ctx context.Context,
	userID int,
	user domain.User,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OptionalTimeOut())
	defer cancel()
	query := `
		UPDATE todoapp.users
		SET
			full_name=$1,
			phone_number=$2,
			version=version+1
		WHERE id=$3 AND version=$4
		RETURNING
			id,
			version,
			full_name,
			phone_number
		`

	row := r.pool.QueryRow(ctx, query, user.FullName, user.PhoneNumber, userID, user.Version)
	var userModel UserModel
	err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.FullName,
		&userModel.PhoneNumber,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, fmt.Errorf(
				"user with id='%v' concurently accessed: %w",
				userID,
				core_errors.ErrConflict,
			)
		}

		return domain.User{}, fmt.Errorf("user scan to patch: %w", err)
	}

	return domain.NewUser(userModel.FullName, userModel.PhoneNumber, userModel.ID, userModel.Version), nil
}
