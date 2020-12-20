package models

//User Model
type User struct {
	ID       	int64 `jsonapi:"primary,users" json:"id"`
	Name 		string `jsonapi:"attr,name" json:"name"`
	Email 		string `jsonapi:"attr,email" json:"email"`
	Password 	string `jsonapi:"attr,password" json:"password"`
}