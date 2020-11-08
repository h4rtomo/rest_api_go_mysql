package helpers


import (
	"os"
	"log"

	"github.com/joho/godotenv"

)
// ResponseData is General response
type ResponseData struct {
	Status  bool    `json:"status"`
	Message string `json:"message"`
	Data    interface{} `json:"data"`
}

//GetENV for get key in .env
func GetENV(key string) string {

	// load .env file
	err := godotenv.Load(".env")
  
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
  
	return os.Getenv(key)
}