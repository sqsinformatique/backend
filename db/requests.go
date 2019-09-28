package db

import (
	"database/sql"
)

func rollbackQuery(query string, args ...interface{}) (rows *sql.Rows, err error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err = tx.Query(query, args...)
	if err != nil {
		return nil, err
	}
	rows.Close()

	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return
}
