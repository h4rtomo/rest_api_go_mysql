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
	
	router.HandleFunc("/user", controllers.HandleAllUsers).Methods("GET")
	router.HandleFunc("/user/store", controllers.HandleCreateUser).Methods("POST")
	router.HandleFunc("/login", controllers.HandleLogin).Methods("POST")

	http.Handle("/", router)
	fmt.Println("Connected to port", port)
	log.Fatal(http.ListenAndServe(":"+port, router))

}