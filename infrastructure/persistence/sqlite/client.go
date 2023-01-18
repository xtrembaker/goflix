package sqlite

import (
	"github.com/jmoiron/sqlx"
	"log"
	"sync"
)

type Client struct {
	db *sqlx.DB
}

func (c *Client) Disconnect() {
	c.db.Close()
}

func (c *Client) IsConnected() bool {
	err := c.db.Ping()
	return err == nil
}

var lock = &sync.Mutex{}
var client *Client

// GetInstance @TODO remove Public
func GetInstance() *Client {
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
		db: db,
	}

	return client
}
