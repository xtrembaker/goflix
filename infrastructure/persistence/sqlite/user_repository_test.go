package sqlite

import (
	"github.com/stretchr/testify/assert"
	"github.com/xtrembaker/goflix/domain"
	"github.com/xtrembaker/goflix/domain/user"
	"testing"
)

func TestFindByUsernameAndPasswordReturnsNotFoundWhenUserDoesNotExists(t *testing.T) {
	userRepository := UserRepository{connection: testingDbConnection}
	u, err := userRepository.FindByUsernameAndPassword("toto", "password")

	assert.Nil(t, u)
	assert.Equal(t, domain.EntityNotFound, err.Error())
}

func TestFindByUsernameAndPasswordReturnsUser(t *testing.T) {
	insertUserRecord()
	userRepository := UserRepository{
		connection: testingDbConnection,
	}
	u, err := userRepository.FindByUsernameAndPassword("gopher", "rocks")
	assert.Nil(t, err)
	assert.Equal(t, &user.User{
		ID:       1,
		Username: "gopher",
		Password: "rocks",
	}, u)
}

func insertUserRecord() {
	testingDbConnection.MustExec("INSERT INTO user (username, password) VALUES (\"gopher\", \"rocks\")")
}
