package main

import (
	"adopet/controllers"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestHomePage(t *testing.T) {
	mockResponse := `{"message":"API rodando com sucesso"}`
	r := SetUpRouter()
	r.GET("/api/v1/home", controllers.GetHome)
	req, _ := http.NewRequest("GET", "/api/v1/home", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}
