package helpers

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	
	"github.com/joho/godotenv"
)

// ResponseData is General response
type ResponseData struct {
	Status  bool    `json:"status"`
	Message string `json:"message"`
	Data    interface{} `json:"data"`
}

// ResponseError is General response
type ResponseError struct {
	Status  bool    `json:"status"`
	Message string `json:"message"`
}

//GetENV for get key .env
func GetENV(key string) string {
	_, b, _, _ := runtime.Caller(0)
	basepath   := filepath.Dir(b)

	// load .env file	
	err := godotenv.Load(basepath + "/../.env")
  
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
  
	return os.Getenv(key)
}