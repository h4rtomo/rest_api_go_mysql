package controllers

import (
	"log"
	"time"
	"net/http"
	"encoding/json"

	models "../models" 
	helpers "../helpers" 

	"github.com/google/jsonapi"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
)

//Claims for JWT
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

var jwtKey = []byte("GOLANGRESTJWT")

//HandleLogin for get JWT
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	type UserBody struct {
		Email 			string `jsonapi:"attr,email" json:"email" validate:"required, email"`
		Password 		string `jsonapi:"attr,password" json:"password"`
	}

	type TokenResponse struct {
		Token 		string `jsonapi:"attr,token" json:"token"`
	}
	var reqbody UserBody
	var users models.User
	var responseToken TokenResponse

	var response helpers.ResponseData
	var responseError helpers.ResponseError

	db := helpers.Connect()
	defer db.Close()

	err := jsonapi.UnmarshalPayload(r.Body, &reqbody)

	if err != nil || reqbody.Email == "" || reqbody.Password == ""  {
		log.Print(err)

		responseError.Status = false
		responseError.Message = "Email and Password Required"

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(responseError)
		return
	} 
	

	rows, err := db.Query("Select id, name, password, email  from users where email = ? ", reqbody.Email)
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&users.ID, &users.Name, &users.Password,  &users.Email); err != nil {
			log.Fatal(err.Error())

		} else {
			
			err := bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(reqbody.Password))
			if err != nil {
				responseError.Status = false
				responseError.Message = "Password not match"

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(200)
				json.NewEncoder(w).Encode(responseError)
				return
			}

			// Declare the expiration time of the token
			// here, we have kept it as 5 minutes
			expirationTime := time.Now().Add(30 * time.Minute)
			// Create the JWT claims, which includes the username and expiry time
			claims := &Claims{
				Email: users.Email,
				StandardClaims: jwt.StandardClaims{
					// In JWT, the expiry time is expressed as unix milliseconds
					ExpiresAt: expirationTime.Unix(),
				},
			}

			// Declare the token with the algorithm used for signing, and the claims
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			// Create the JWT string
			tokenString, err := token.SignedString(jwtKey)
			if err != nil {
				// If there is an error in creating the JWT return an internal server error
				log.Println(err)
			}
			responseToken.Token = tokenString
		}
	}

	response.Status = true
	response.Message = "Data Token"
	response.Data = responseToken

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(response)
}