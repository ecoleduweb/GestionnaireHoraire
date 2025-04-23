package controllers

import (
	"llio-api/models/DTOs"
	"llio-api/services"
	"net/http"

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

	c.JSON(http.StatusOK, gin.H{"firstName": user.FirstName, "lastName": user.LastName})
}
