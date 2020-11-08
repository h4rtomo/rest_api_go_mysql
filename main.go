package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/google/jsonapi"

	controllers "./controllers" 
	helpers "./helpers"
)

func main() {
	port := helpers.GetENV("PORT")

	router := mux.NewRouter()
	router.HandleFunc("/users", controllers.HandleAllUsers).Methods("GET")
	http.Handle("/", router)
	fmt.Printf("Connected to port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))

}