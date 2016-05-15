package handler

import (
	"testing"
	"fmt"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"common"
)

func TestGetAllProducts(t *testing.T) {
	router := common.GetTestModeGinRouter();
	handler := GetAllProducts
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
	router := common.GetTestModeGinRouter();
	router.GET("/search", FindProduct)
	req, err := http.NewRequest("GET", "/search?query=5s", nil)
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	fmt.Println(resp.Body)
	assert.Equal(t, resp.Code, 200)
}


func TestFindProducts_WithEmptySearchString(t *testing.T){
	router := common.GetTestModeGinRouter();
	router.GET("/search", FindProduct)
	req, err := http.NewRequest("GET", "/search?query=", nil)
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	fmt.Println(resp.Body)
	assert.Equal(t, resp.Code, 200)
}