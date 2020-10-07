package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func CreatConnection() *sql.DB {

	db, err := sql.Open("mysql", "root:my-pw@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Errorf("No DB conn")
	}
	err = db.Ping()
	if err != nil {

	}
	return db
}
