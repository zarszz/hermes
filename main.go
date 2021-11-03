package main

import (
	"fmt"
	"hermes/controllers"
	"hermes/models/sql"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main()  {
	err := godotenv.Load()
	if err != nil {
		msg := fmt.Sprintf("An error occured : %s", err.Error())
		panic(msg)		
	}
	conn, err := sql.NewSQLInstance()
	if err != nil {
		msg := fmt.Sprintf("An error occured : %s", err.Error())
		panic(msg)
	}
	productController := controllers.NewProduct(conn)
	if err != nil {
		msg := fmt.Sprintf("An error occured : %s", err.Error())
		panic(msg)
	}

	r := gin.Default()

	r.GET("/products", productController.Get)
	r.GET("/products/:id", productController.GetByPK)
	r.POST("/products", productController.Create)
	r.PUT("/products/:id", productController.Update)
	r.DELETE("products/:id", productController.DeleteByPK)
	
	r.Run(":8080")
}
