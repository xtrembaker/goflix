package sqlite

import (
	"github.com/jmoiron/sqlx"
	"log"
)

var connection *sqlx.DB

func Connect() *sqlx.DB {
	if connection != nil {
		return connection
	}
	db, err := sqlx.Connect("sqlite3", "goflix.db")
	if err != nil {
		log.Fatal(err)
	}
	connection = db
	return connection
}

func Disconnect() {
	connection.Close()
}

func IsConnected() bool {
	err := connection.Ping()
	return err == nil
}
