package sqlite

import (
	"github.com/jmoiron/sqlx"
	"log"
)

var schema = `
CREATE TABLE IF NOT EXISTS movie
(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT,
	release_date TEXT,
	duration INTEGER,
	trailer_url TEXT
)
`

// @TODO That would be great to be able to launch that in CLI (add a flag to do so)
func ExecMigration(db *sqlx.DB) {
	db.MustExec(schema)
	log.Println("Migration executed")
}
