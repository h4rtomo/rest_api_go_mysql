package helpers

import (
	"fmt"
	"database/sql"
	"log"
)

//Connect to Database
func Connect() *sql.DB {
	user := GetENV("DB_USER")
	password := GetENV("DB_PASSWORD")
	host := GetENV("DB_HOST")
	port := GetENV("DB_PORT")
	dbName := GetENV("DB_NAME")

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName)
	db, err := sql.Open("mysql", connection)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
