package users_postgres_repository

import (
	"context"
	"fmt"

	"github.com/MrLaplace19/ToDoList/internal/core/domain"
)

func (r *UsersRepository) GetUser(
		ctx context.Context,
		limit *int,
		offset *int,
	)([]domain.User,error){
		ctx, cancel := context.WithTimeout(ctx, r.pool.OptionalTimeOut())
		defer cancel()

		query := `
		SELECT id,version,full_name, phone_number
		FROM todoapp.users
		ORDER BY id ASC
		LIMIT $1
		OFFSET $2
		`

		rows, err := r.pool.Query(ctx, query, limit, offset)
		if err != nil {
			return nil, fmt.Errorf("select users: %w", err)
		}
		defer rows.Close()

		var userModel []UserModel
		for rows.Next(){
			var usermodel UserModel
			err := rows.Scan(
				&usermodel.ID,
				&usermodel.Version,
				&usermodel.FullName,
				&usermodel.PhoneNumber,
			)
			if err != nil {
				return nil, fmt.Errorf("scan users: %w", err)
			}
			userModel = append(userModel, usermodel)
		}
		if err := rows.Err(); err != nil {
			return nil, fmt.Errorf("next rows: %w", err)
		}

		userDomains := UserDomainsFromModel(userModel)
		return userDomains, nil
	}