package controllers

import (
	"log"
	"net/http"
	"encoding/json"

	models "../models" 
	helpers "../helpers" 

	"github.com/google/jsonapi"
	"golang.org/x/crypto/bcrypt"
)

//HandleAllUsers for all user
func HandleAllUsers(w http.ResponseWriter, r *http.Request) {
	var users models.User
	var arrUser []models.User
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

//HandleCreateUser for create new user
func HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	type UserBody struct {
		ID       	int64 `jsonapi:"primary,users" json:"id"`
		Email 			string `jsonapi:"attr,email" json:"email" validate:"required, email"`
		Password 		string `jsonapi:"attr,password" json:"password"`
		Name 		string `jsonapi:"attr,name" json:"name"`
	}
	var user UserBody
	var response helpers.ResponseData

	db := helpers.Connect()
	defer db.Close()

	err := jsonapi.UnmarshalPayload(r.Body, &user)

	if err != nil {
		log.Print(err)
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)

	if err != nil {
		log.Print(err)
		return
	}

	query, err := db.Prepare("INSERT INTO users (name, email, password) VALUES (?, ?, ?)")
	
	if err != nil {
		log.Print(err)
		return
	}

	result, err := query.Exec(user.Name, user.Email, hash)

	if err != nil {
		log.Print(err)
		return
	}

	lastID, err := result.LastInsertId()
	user.Password = ""

	if err != nil {
		log.Print(err)
		return
	}

	user.ID = lastID

	response.Status = true
	response.Message = "User created"
	response.Data = &user

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(response)
}