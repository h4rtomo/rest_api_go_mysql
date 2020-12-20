package test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	controllers "../controllers" 
	helpers "../helpers" 

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	
	"github.com/stretchr/testify/assert"
	
	_ "github.com/go-sql-driver/mysql"
)

func Test_HandleAllUsers(t *testing.T){
	request, _ := http.NewRequest("GET", "/users", nil)

	response := httptest.NewRecorder()

	controllers.HandleAllUsers(response, request)
	
	t.Log("==================== Start ==============")
	t.Log(response.Body)
	t.Log("==================== End ==============\n")

	body, _ := ioutil.ReadAll(response.Body)

	res := helpers.ResponseData{}

	json.Unmarshal([]byte(body), &res)
	t.Log("==================== Start ==============")
	t.Log(res.Message)
	t.Log("==================== End ==============\n")

	assert.Equal(t, 200, response.Code, "Invalid Rsponse Code")
	assert.Equal(t, "Data fetched", res.Message, "Invalid Message ")
}

func Test_HandleCreateUser(t *testing.T){
	name := "Rudi H"
	email := "rudi@gmail.com"

	db, mock, err := sqlmock.New()

	if err != nil {
		t.Log("==================== Start ==============")
		t.Log(err)
		t.Log("==================== End ==============\n")
	}
	defer db.Close()

	mock.ExpectPrepare("INSERT INTO users (name, email) VALUES (?, ?)")
	mock.ExpectExec("INSERT INTO users (name, email) VALUES (?, ?)").WithArgs(name, email).WillReturnResult(sqlmock.NewResult(10,1))

	data := map[string]interface{}{
		"data":   map[string]interface{}{
			"attributes":  map[string]interface{}{
				"name": name,
				"email": email,
			},
		},
	}

	requestBody, _ := json.Marshal(data)
	request, _ := http.NewRequest("POST", "/user/store", bytes.NewBuffer(requestBody))

	response := httptest.NewRecorder()

	controllers.HandleCreateUser(response, request)
	
	t.Log("==================== Start ==============")
	t.Log(response.Body)
	t.Log("==================== End ==============\n")

	body, _ := ioutil.ReadAll(response.Body)

	res := helpers.ResponseData{}

	json.Unmarshal([]byte(body), &res)
	t.Log("==================== Start ==============")
	t.Log(res.Message)
	t.Log("==================== End ==============\n")

	assert.Equal(t, 200, response.Code, "Invalid Rsponse Code")
	assert.Equal(t, "User created", res.Message, "Invalid Message ")
}
