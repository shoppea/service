package handler

import (
	"testing"
	"fmt"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"common"
)

func TestGetUserCart(t *testing.T) {
	router := common.GetTestModeGinRouter();
	router.GET("/cart/:id", GetUserCart)
	req, err := http.NewRequest("GET", "/cart/6", nil)
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	fmt.Println(resp.Body)
	assert.Equal(t, resp.Code, 200)
}
