package main

import (
	"db"
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"handler"
	"net/http"
	"path/filepath"
	"github.com/itsjamie/gin-cors"
	"time"
)

func init() {
	// Create database connection pool with mysql db
	dbPool := db.SharedConnection()
	// Any error would simply stops execution
	if dbPool.Error != nil {
		logrus.Error(dbPool.Error)
		panic("failed to initialse database")
	}
}

func main() {
	r := gin.Default()

	r.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET, PUT, POST, DELETE",
		RequestHeaders: "Origin, Authorization, Content-Type",
		ExposedHeaders: "",
		MaxAge: 50 * time.Second,
		Credentials: true,
		ValidateHeaders: false,
	}))

	// Serve static swagger pages for api docs
	r.StaticFile("/swagger_api","./src/main/swagger.json")
	r.StaticFS("/docs",http.Dir("./dist"))

	r.GET("/fs", func(c *gin.Context) {
		files, _ := filepath.Glob("*")
		c.JSON(200,files)
	})

	// swagger:route GET /categories category
	//
	// Lists all available categories
	//
	// This will show all available main categories by default.
	// You can get subcategories in /subcategory api
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Responses:
	//       default: genericError
	//       200: someResponse
	//       422: validationError
	r.GET("/categories", handler.GetAllCategories)

	// swagger:route POST /category category
	//
	// Create new category
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Responses:
	//       default: genericError
	//       200: someResponse
	//       422: validationError
	r.POST("/category", handler.CreateCategory)

	// swagger:route GET /genders gender
	//
	// Lists all available genders
	//
	// This will show all available genders by default.
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Responses:
	//       default: genericError
	//       200: someResponse
	//       422: validationError
	r.GET("/genders", handler.GetAllGenders)

	// swagger:route POST /gender gender
	//
	// Create new gender

	// We added all available genders of planet earth,
	// so please don't play with API ;)
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Responses:
	//       default: genericError
	//       200: someResponse
	//       422: validationError
	r.POST("/gender", handler.CreateGender)

	// swagger:route POST /sub_category category
	//
	// Create new Sub Category
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Responses:
	//       default: genericError
	//       200: someResponse
	//       422: validationError
	r.POST("/sub_category", handler.CreateSubCategory)

	// swagger:route POST /product product
	//
	// Create new Sub Product
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Responses:
	//       default: genericError
	//       200: someResponse
	//       422: validationError
	r.POST("/product", handler.CreateProduct)

	// swagger:route POST /products product
	//
	// Create new Sub Product
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Responses:
	//       default: genericError
	//       200: someResponse
	//       422: validationError
	r.GET("/products", handler.GetAllProducts)

	r.POST("/user", handler.CreateUser)

	r.GET("/help", func(c *gin.Context) {
		c.JSON(http.StatusOK, r.Routes())
	})

	// swagger:route GET /search product
	//
	// Search products with name containing search string
	//
	//     Consumes:
	//     -
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Responses:
	//       default: genericError
	//       200: someResponse
	//       422: validationError
	r.GET("/search", handler.FindProduct)

	r.Run(":8081")
}