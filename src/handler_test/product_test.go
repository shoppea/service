package handler_test

import (
	"testing"
	"fmt"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"github.com/gin-gonic/gin"
	"handler"
)

func TestGetllProducts(t *testing.T) {

	gin.SetMode(gin.TestMode);
	router := gin.Default();
	handler := handler.GetAllProducts
	router.GET("/products", handler)
	req, err := http.NewRequest("GET", "/products", nil)
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	fmt.Println(resp.Body)
	assert.Equal(t, resp.Code, 200)
}

func TestFindProducts(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/search", handler.FindProduct)
	req, err := http.NewRequest("GET", "/search?query=5s", nil)
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	fmt.Println(resp.Body)
	assert.Equal(t, resp.Code, 200)
}
