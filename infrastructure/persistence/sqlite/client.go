package sqlite

import (
	"github.com/jmoiron/sqlx"
	"log"
	"sync"
)

type Client struct {
	connection *sqlx.DB
}

func (c *Client) Disconnect() {
	c.connection.Close()
}

func (c *Client) IsConnected() bool {
	err := c.connection.Ping()
	return err == nil
}

var lock = &sync.Mutex{}
var client *Client

func getInstance() *Client {
	if client != nil {
		return client
	}
	lock.Lock()
	defer lock.Unlock()

	db, err := sqlx.Connect("sqlite3", "goflix.db")
	if err != nil {
		log.Fatal(err)
	}

	client = &Client{
		connection: db,
	}

	return client
}
