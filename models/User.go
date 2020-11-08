package models

//Users Model
type Users struct {
	ID       	int `jsonapi:"primary,users" json:"id"`
	Name 		string `jsonapi:"attr,name" json:"name"`
}
