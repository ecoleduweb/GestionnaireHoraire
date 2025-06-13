package controllers

import (
	"llio-api/models/DTOs"
	"llio-api/models/enums"
	"llio-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var userSTR = "utilisateur"

func GetUserInfo(c *gin.Context) {
	userInfo, isExist := c.Get("current_user")
	if !isExist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Utilisateur non trouvé dans le contexte"})
		return
	}

	currentUser, ok := userInfo.(*DTOs.UserDTO)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Type d'utilisateur invalide"})
		return
	}

	user, err := services.GetUserByEmail(currentUser.Email)
	if err != nil {
		handleError(c, err, userSTR)
		return
	}

	c.JSON(http.StatusOK, gin.H{"firstName": user.FirstName, "lastName": user.LastName, "role": user.Role})
}

func GetAllUsers(c *gin.Context) {
	userRoleStr := c.Query("role")

	var userRole *enums.UserRole

	if userRoleStr != "" {
		roleValue, err := strconv.Atoi(userRoleStr)
		if err != nil || roleValue < 0 || roleValue > 2 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Role invalide"})
			return
		}
		role := enums.UserRole(roleValue)
		userRole = &role
	}

	users, err := services.GetAllUsers(userRole)
	if err != nil {
		handleError(c, err, userSTR)
		return
	}

	if users == nil {
		users = []*DTOs.UserDTO{}
	}

	c.JSON(http.StatusOK, users)
}

func GetManagerAdminUsers(c *gin.Context) {
	role1 := enums.UserRole(1)
	users1, err := services.GetAllUsers(&role1)
	if err != nil {
		handleError(c, err, userSTR)
		return
	}

	role2 := enums.UserRole(2)
	users2, err := services.GetAllUsers(&role2)
	if err != nil {
		handleError(c, err, userSTR)
		return
	}

	var allUsers []*DTOs.UserDTO
	if users1 != nil {
		allUsers = append(allUsers, users1...)
	}
	if users2 != nil {
		allUsers = append(allUsers, users2...)
	}

	if allUsers == nil {
		allUsers = []*DTOs.UserDTO{}
	}

	c.JSON(http.StatusOK, allUsers)
}

func UpdateUserRole(c *gin.Context) {
	// Get the current user from context
	userInfo, isExist := c.Get("current_user")
	if !isExist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Utilisateur non trouvé dans le contexte"})
		return
	}

	// Get the user ID from the URL parameters
	userID := c.Param("id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de l'utilisateur manquant"})
		return
	}

	userID_int, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID d'utilisateur invalide"})
		return
	}

	if userID_int == userInfo.(*DTOs.UserDTO).Id {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Vous ne pouvez pas changer votre propre rôle"})
		return
	}

	// Parse the role from the request body
	var roleRequest struct {
		Role *int `json:"role" binding:"required"`
	}

	if err := c.ShouldBindJSON(&roleRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Rôle invalide: " + err.Error()})
		return
	}

	// Import the enums package and convert string to UserRole
	userRole, err := enums.ParseUserRole(*roleRequest.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Rôle invalide: " + err.Error()})
		return
	}

	userDTO := DTOs.UserDTO{
		Id:   userID_int,
		Role: userRole,
	}

	// Update the user's role
	user, err := services.UpdateUserRole(&userDTO)
	if err != nil {
		handleError(c, err, userSTR)
		return
	}

	c.JSON(http.StatusOK, user)
}
