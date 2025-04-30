package services

import (
	"llio-api/models/DAOs"
	"llio-api/models/DTOs"
	"llio-api/repositories"

	"github.com/jinzhu/copier"
)

func FirstOrCreateUser(userDTO *DTOs.UserDTO) (*DTOs.UserDTO, error) {

	user := &DAOs.User{}
	err := copier.Copy(user, userDTO)
	if err != nil {
		return nil, err
	}
	userDAOAdded, err := repositories.FirstOrCreateUser(user)
	if err != nil {
		return nil, err
	}

	userDTOResponse := &DTOs.UserDTO{}
	err = copier.Copy(userDTOResponse, userDAOAdded)
	return userDTOResponse, err
}

func GetUserById(id string) (*DTOs.UserDTO, error) {
	user, err := repositories.GetUserById(id)
	if err != nil {
		return nil, err
	}

	userDTO := &DTOs.UserDTO{}
	err = copier.Copy(userDTO, user)

	return userDTO, err
}

func GetUserByEmail(email string) (*DTOs.UserDTO, error) {
	user, err := repositories.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	userDTO := &DTOs.UserDTO{}
	err = copier.Copy(userDTO, user)

	return userDTO, err
}

func GetAllUsers() ([]*DTOs.UserDTO, error) {
	users, err := repositories.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var usersDTOs []*DTOs.UserDTO
	for _, user := range users {
		userDTO := &DTOs.UserDTO{}
		err = copier.Copy(userDTO, user)
		usersDTOs = append(usersDTOs, userDTO)
	}

	return usersDTOs, err
}
