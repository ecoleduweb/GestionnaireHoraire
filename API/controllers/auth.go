package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

// Function pour l'authentification des utilisateurs
func Auth(c *gin.Context) {
	provider := c.Param("provider")
	q := c.Request.URL.Query()
	q.Add("provider", provider)
	c.Request.URL.RawQuery = q.Encode()
	if gothUser, err := gothic.CompleteUserAuth(c.Writer, c.Request); err == nil {
		fmt.Println(gothUser)
	} else {
		gothic.BeginAuthHandler(c.Writer, c.Request)
	}
}

// Function pour la récupération des informations de l'utilisateur (callback après authentification)
func GetAuthCallback(c *gin.Context) {
	provider := c.Param("provider")
	q := c.Request.URL.Query()
	q.Add("provider", provider)
	c.Request.URL.RawQuery = q.Encode()
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Erreur lors de l'authentification", "message": err.Error()})
		return
	}
	fmt.Println(user.Email)
	// Redirection vers une autre page si nécessaire
}

// Function pour la déconnexion des utilisateurs
func Logout(c *gin.Context) {
	gothic.Logout(c.Writer, c.Request)
	c.Writer.Header().Set("Location", "/")
	c.Writer.WriteHeader(http.StatusTemporaryRedirect)
}
