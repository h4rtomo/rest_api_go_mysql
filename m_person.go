package main

//Users Model
type Users struct {
	ID        string `form:"id" json:"id"`
	FirstName string `form:"firstname" json:"firstname"`
	LastName  string `form:"lastname" json:"lastname"`
}

// ResponseData is General response
type ResponseData struct {
	Status  bool    `json:"status"`
	Message string `json:"message"`
	Data    interface{} `json:"data"`
}
