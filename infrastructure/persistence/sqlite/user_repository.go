package sqlite

import (
	goErrors "errors"
	"github.com/jmoiron/sqlx"
	"github.com/xtrembaker/goflix/domain"
	"github.com/xtrembaker/goflix/domain/user"
)

type UserRepository struct {
	connection *sqlx.DB
}

func (r UserRepository) FindByUsernameAndPassword(username, password string) (*user.User, error) {
	u := &user.User{}
	err := r.connection.Get(u, "SELECT * FROM user WHERE username=$1 and password=$2", username, password)
	if err != nil {
		return nil, goErrors.New(domain.EntityNotFound)
	}
	return u, nil
}

func UserRepositoryFactory() user.UserRepository {
	return &UserRepository{
		connection: Connect(),
	}
}
