package handler

import (
	"testing"
	"common"
	"fmt"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"bytes"
)

func TestAddShopper(t *testing.T) {
	r := common.GetTestModeGinRouter()
	r.POST("/shopper/new",AddShopper)
	shopper := []byte(`
	{
 		"name": "Abc Corps",
 		"description": "Selling electronics materials",
 		"address": "A-301,Shyamal Apartment,Pune-48",
 		"vat_no": "V712718820",
 		"contact_no": 7276,
 		"latitude":"18.520430",
 		"longitude":"73.856744",
 		"area_name": "Kondhawa-BK",
 		"id": 14
	}
	`)
	req, err := http.NewRequest("POST", "/shopper/new", bytes.NewBuffer(shopper))
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	fmt.Println(resp.Body)
	assert.Equal(t, resp.Code, 200)
}


func TestAddShopperNegative(t *testing.T) {
	r := common.GetTestModeGinRouter()
	r.POST("/shopper/new",AddShopper)
	shopper := []byte(`
	{
 		"namea": "Abc Corps",
 		"description": "Selling electronics materials",
 		"address": "A-301,Shyamal Apartment,Pune-48",
 		"vat_no": "V712718820",
 		"contact_no": 7276,
 		"latitude":"18.520430",
 		"longitude":"73.856744",
 		"area_name": "Kondhawa-BK",
 		"id": 14
	}
	`)
	req, err := http.NewRequest("POST", "/shopper/new", bytes.NewBuffer(shopper))
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	fmt.Println(resp.Body)
	assert.Equal(t, resp.Code,400)
}
