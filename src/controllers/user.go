package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ladmakhinima/go-simple-crud-api/src/models"
	"github.com/ladmakhinima/go-simple-crud-api/src/services"
)

func SaveController(ctx *gin.Context) {
	var requestBody models.User
	ctx.BindJSON(&requestBody)
	user := services.Save(requestBody)
	ctx.JSON(201, gin.H{
		"user": user,
	})
}

func FindByFirstName(ctx *gin.Context) {
	FirstName := ctx.Param("FirstName")
	user := services.FindByFirstName(FirstName)
	ctx.JSON(200, gin.H{
		"user": user,
	})
}

func FindAllUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"users": services.UsersData,
	})
}
