package handler

import (
	"testing"
	"fmt"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"common"
)

func TestGetAllCategories(t *testing.T) {
	router := common.GetTestModeGinRouter();
	router.GET("/categories", GetAllCategories)
	req, err := http.NewRequest("GET", "/categories", nil)
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	fmt.Println(resp.Body)
	assert.Equal(t, resp.Code, 200)
}