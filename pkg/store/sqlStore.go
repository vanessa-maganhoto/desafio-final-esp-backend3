package store

import (
	"database/sql"
)

type SqlStore struct {
	db *sql.DB
}

