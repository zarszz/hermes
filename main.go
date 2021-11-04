package main

import (
	"fmt"
	"hermes/controllers"
	"hermes/middlewares"
	"hermes/models/sql"
	"os"
	"path/filepath"

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
	userController := controllers.NewUserController(conn)

	if err != nil {
		msg := fmt.Sprintf("An error occured : %s", err.Error())
		panic(msg)
	}

	r := gin.Default()
	
	cwd, _ := os.Getwd()
	imagesPath := filepath.Join(cwd, "images")
	if _, err := os.Stat(imagesPath); err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(imagesPath, 0755)
		}
	}
	r.Static("/images", filepath.Join(cwd, "images"))
	
	r.POST("/users", userController.Register)
	r.POST("/users/login", userController.Login)

	authorized := r.Group("/")
	authorized.Use(middlewares.AuthMiddleware())
	{
		authorized.GET("/products", productController.Get)
		authorized.GET("/products/:id", productController.GetByPK)
		authorized.POST("/products", productController.Create)
		authorized.PUT("/products/:id", productController.Update)
		authorized.DELETE("products/:id", productController.DeleteByPK)
	
		authorized.GET("/users", userController.Get)
		authorized.GET("users/:id", userController.GetByPK)
		authorized.PUT("/users/:id", userController.Update)
		authorized.DELETE("users/:id", userController.DeleteByPK)
	}	
	
	r.Run(":" + os.Getenv("LISTEN_PORT"))
}
