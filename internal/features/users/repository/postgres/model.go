package users_postgres_repository

import "github.com/MrLaplace19/ToDoList/internal/core/domain"

type UserModel struct {
	ID          int
	Version     int
	FullName    string
	PhoneNumber *string
}

func UserDomainsFromModel(users []UserModel) []domain.User {
	userDomains := make([]domain.User, len(users))
	for i, v := range users {
		userDomains[i] = domain.NewUser(
			v.FullName,
			v.PhoneNumber,
			v.ID,
			v.Version,
		)
	}
	return userDomains
}
