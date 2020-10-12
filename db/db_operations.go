package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func CreatConnection() (*sql.DB , error) {

	db, err := sql.Open("mysql", "root:my-pw@tcp(127.0.0.1:3306)/test")
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db , nil
}
