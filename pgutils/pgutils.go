package pgutils

import (
	"database/sql"

	"github.com/lib/pq"
)

func Open(dsn string) (*sql.DB, error) {

	// parse url
	c, err := pq.ParseURL(dsn)
	if err != nil {
		return nil, err
	}

	// create connection
	db, err := sql.Open("postgres", c)
	if err != nil {
		return nil, err
	}

	// test connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// return
	return db, nil
}
