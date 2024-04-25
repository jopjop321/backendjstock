package server

import (
	"database/sql"
)

func CreateConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:jobben321@tcp(127.0.0.1:3306)/data_jstock")
	if err != nil {
		return nil, err
	}
	return db, nil
}