package pg

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const dbVersion = "20221116142749"

func NewDB(user, password, addr, port, dbname string) (*sql.DB, error) {
	db, err := sql.Open("postgres",
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			user, password, addr, port, dbname))
	if err != nil {
		return nil, err
	}

	row := db.QueryRow(`SELECT version_id FROM goose_db_version WHERE is_applied = TRUE ORDER BY version_id DESC LIMIT 1`)
	if row.Err() != nil {
		return nil, fmt.Errorf("cannot get goose version: %w", row.Err())
	}

	var version string
	err = row.Scan(&version)
	if err != nil {
		return nil, fmt.Errorf("cannot scan goose version: %w", err)
	}

	if version != dbVersion {
		return nil, fmt.Errorf("db version mismatch: current=%s; required=%s", version, dbVersion)
	}

	return db, nil
}
