package repositories

import (
	"llio-api/database"
	"llio-api/models/DAOs"
)

func CreateUser(user *DAOs.User) (*DAOs.User, error) {
	err := database.DB.Create(user).Error
	return user, DBErrorManager(err)
}

func GetUserById(id string) (*DAOs.User, error) {
	var user DAOs.User

	err := database.DB.First(&user, id).Error
	return &user, DBErrorManager(err)
}

func GetUserByEmail(email string) (*DAOs.User, error) {
	var user DAOs.User

	err := database.DB.Where("email = ?", email).First(&user).Error
	return &user, DBErrorManager(err)
}
