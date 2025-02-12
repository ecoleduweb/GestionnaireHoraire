package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Exemple de données
var users = []string{"Alice", "Bob", "Charlie"}

// Récupérer tous les utilisateurs
func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"users": users})
}

// Ajouter un utilisateur
func AddUser(c *gin.Context) {
	var newUser struct {
		Name string `json:"name"`
	}
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	users = append(users, newUser.Name)
	c.JSON(http.StatusOK, gin.H{"message": "User added", "users": users})
}
