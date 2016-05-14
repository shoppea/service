package handler

import (
	"testing"
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"db"
	"product"
	"strings"
)

//*********************************************
//TestMain will get called For every test you run
//-------Note: TODO: need to place in appropriate file
func TestMain(m *testing.M) {
	_Init();
	result := m.Run();
	_TearDown();
	os.Exit(result);
}

func _Init() {
	fmt.Println("Setting up test...\n");

}

func _TearDown() {
	fmt.Println("Tearing down tests...\n");
	fmt.Println("All testing stubs deleted successfully");
}
//-*******************************************


func TestGetAllGenders(t *testing.T) {
	gin.SetMode(gin.TestMode);
	router := gin.Default();
	router.GET("/genders", GetAllGenders)
	req, err := http.NewRequest("GET", "/genders", nil)
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	fmt.Println(resp.Body)
	assert.Equal(t, resp.Code, 200)
}

func TestCreateGender(t *testing.T) {
	gin.SetMode(gin.TestMode);
	router := gin.Default();
	router.POST("/gender", CreateGender);

	testGenderJson := `{"id":9,"name":"TestGender"}`;
	reader := strings.NewReader(testGenderJson)

	req, err := http.NewRequest("POST", "/gender", reader);
	if err != nil {
		fmt.Println(err);
	}
	resp := httptest.NewRecorder();
	router.ServeHTTP(resp, req);
	dbpool := db.SharedConnection();
	var testGender product.Gender;
	err = dbpool.Find(&testGender, product.Gender{Id: 9}).Error;
	if err != nil {
		fmt.Println(err);
	}

	//*********************************************
	//-----------Deleting test gender-------------
	//*********************************************
	err = dbpool.Delete(&testGender, product.Gender{Id:9}).Error;
	if err != nil {
		fmt.Println(err);
	}
	//----------------------------------------------

	assert.Equal(t, resp.Code == 200 && testGender.Name == "TestGender", true);
}
