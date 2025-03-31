package repository

import (
	"ticketsbooking/libs/entities"
	"ticketsbooking/services/auth/internal/interfaces"

	"github.com/pkg/errors"
)

type postgresRepo struct {
	executor interfaces.QueryExecutor
}

func New(executor interfaces.QueryExecutor) interfaces.Repository {
	return &postgresRepo{executor: executor}
}

func (repo *postgresRepo) CreateUser(user entities.User) error {
	query := `INSERT INTO users
	(name,email,password)
	VALUES (:name,:email,:password)`
	_, err := repo.executor.NamedExec(query, user)
	return errors.Wrap(err, "calling sqlx NamedQuery")
}

func (repo *postgresRepo) GetUserByEmail(user entities.User) (entities.User, error) {
	foundedUser := entities.User{}
	err := repo.executor.Get(&foundedUser, "SELECT * FROM users WHERE email = $1", user.Email)
	return foundedUser, errors.Wrap(err, "calling sqlx Get")
}
