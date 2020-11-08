package main

import (
	"database/sql"
	"log"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:Rudi1234@tcp(localhost:3306)/rest_go")

	if err != nil {
		log.Fatal(err)
	}

	return db
}
