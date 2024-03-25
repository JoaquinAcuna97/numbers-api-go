package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/numbers", getNumbers)
	router.POST("/numbers", postNumbers)
	router.GET("/numbers/:number", getCategoryOfNumber)

	return router
}

func TestGetNumbers(t *testing.T) {
	router := setupRouter()
	numbers = []number{{Number: 3, Category: "Type 1"}, {Number: 5, Category: "Type 2"}}
	req, _ := http.NewRequest("GET", "/numbers", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var result []number
	err := json.Unmarshal(resp.Body.Bytes(), &result)
	assert.NoError(t, err)
	assert.Equal(t, numbers, result)
}

func TestPostNumbers(t *testing.T) {
	router := setupRouter()
	input := number{Number: 15}
	reqBody, _ := json.Marshal(input)
	req, _ := http.NewRequest("POST", "/numbers", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)

	var result []number
	err := json.Unmarshal(resp.Body.Bytes(), &result)
	expectedResult := number{Number: 15, Category: "Type 3"}
	assert.NoError(t, err)
	assert.Contains(t, result, expectedResult)
}

func TestGetCategoryOfNumber(t *testing.T) {
	router := setupRouter()
	numbers = []number{{Number: 3, Category: "Type 1"}, {Number: 5, Category: "Type 2"}}
	req, _ := http.NewRequest("GET", "/numbers/5", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var result number
	err := json.Unmarshal(resp.Body.Bytes(), &result)
	assert.NoError(t, err)
	assert.Equal(t, numbers[1], result)
}

func TestInvalidGetCategoryOfNumber(t *testing.T) {
	router := setupRouter()
	req, _ := http.NewRequest("GET", "/numbers/abc", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestNotFoundGetCategoryOfNumber(t *testing.T) {
	router := setupRouter()
	req, _ := http.NewRequest("GET", "/numbers/10", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestPostNumbersInvalidInput(t *testing.T) {
	router := setupRouter()
	reqBody := []byte(`{"num": 10}`)
	req, _ := http.NewRequest("POST", "/numbers", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}
