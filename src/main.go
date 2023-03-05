package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ladmakhinima/go-simple-crud-api/src/controllers"
)

func main() {
	s := gin.Default()
	s.GET("/", controllers.SaveController)
	s.GET("/:FirstName", controllers.FindByFirstName)
	s.POST("/", controllers.FindAllUsers)
	s.Run(":3000")
}
