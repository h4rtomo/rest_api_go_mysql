package test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	controllers "../controllers" 
	helpers "../helpers" 

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