package controllers

import (
	"log"
	"net/http"
	"encoding/json"

	models "../models" 
	helpers "../helpers" 
)

//HandleAllUsers for all user
func HandleAllUsers(w http.ResponseWriter, r *http.Request) {
	var users models.Users
	var arrUser []models.Users
	var response helpers.ResponseData
	//var data struct{}

	db := helpers.Connect()
	defer db.Close()

	rows, err := db.Query("Select id, name from users")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&users.ID, &users.Name); err != nil {
			log.Fatal(err.Error())

		} else {
			arrUser = append(arrUser, users)
		}
	}

	response.Status = true
	response.Message = "Data fetched"
	response.Data = arrUser

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(response)

}


