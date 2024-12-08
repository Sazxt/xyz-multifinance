package test

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"xyz-multifinance/config"
	"xyz-multifinance/controllers"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	config.ConnectDatabase()
	router := gin.Default()
	router.GET("/users", controllers.GetUsers)

	req, _ := http.NewRequest("GET", "/users", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var response map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &response)

	log.Println(response)

	users, ok := response["users"].([]interface{})
	assert.True(t, ok, "Expected 'users' key in response")
	assert.NotEmpty(t, users, "Expected non-empty 'users' array")

	firstUser := users[0].(map[string]interface{})
	assert.Equal(t, float64(1), firstUser["ID"])
	assert.Equal(t, "Aldi", firstUser["FullName"])
	assert.Equal(t, "Tanggerang", firstUser["BirthPlace"])
	assert.Equal(t, "16-03-2000", firstUser["BirthDate"])

	loanLimits, ok := firstUser["LoanLimits"].([]interface{})
	assert.True(t, ok, "Expected 'LoanLimits' key in user data")
	assert.NotEmpty(t, loanLimits, "Expected non-empty 'LoanLimits' array")

	// firstLoanLimit := loanLimits[0].(map[string]interface{})
	// assert.Equal(t, float64(16), firstLoanLimit["ID"])
	// assert.Equal(t, float64(1), firstLoanLimit["Tenor"])
	// assert.Equal(t, float64(13000), firstLoanLimit["Amount"])
}
