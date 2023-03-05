package main

import "github.com/gin-gonic/gin"

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var users []User

type UserServices interface {
	Save() User
	FindAll() []User
	FindByFirstName(FirstName string) User
}

func (u User) Save() User {
	users = append(users, u)
	return u
}

func FindAll() []User {
	return users
}

func FindByFirstName(FirstName string) User {
	var user User
	for _, u := range users {
		if u.FirstName == FirstName {
			user = u
		}
	}
	return user
}

func main() {
	s := gin.Default()
	s.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"users": users,
		})
	})
	s.GET("/:FirstName", func(ctx *gin.Context) {
		FirstName := ctx.Param("FirstName")
		user := FindByFirstName(FirstName)
		ctx.JSON(200, gin.H{
			"user": user,
		})
	})
	s.POST("/", func(ctx *gin.Context) {
		var requestBody User
		ctx.BindJSON(&requestBody)
		user := requestBody.Save()
		ctx.JSON(201, gin.H{
			"user": user,
		})
	})
	s.Run(":3000")
}
