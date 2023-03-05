package services

import "github.com/ladmakhinima/go-simple-crud-api/src/models"

var UsersData []models.User

type UserServices interface {
	Save() models.User
	FindAll() []models.User
	FindByFirstName(FirstName string) models.User
}

func Save(u models.User) models.User {
	UsersData = append(UsersData, u)
	return u
}

func FindAll() []models.User {
	return UsersData
}

func FindByFirstName(FirstName string) models.User {
	var user models.User
	for _, u := range UsersData {
		if u.FirstName == FirstName {
			user = u
		}
	}
	return user
}
