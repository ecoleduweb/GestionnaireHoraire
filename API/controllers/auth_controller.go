package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func Auth(c *gin.Context) {
	provider := c.Param("provider")
	log.Printf("Tentative d'authentification avec le provider %s", provider)
	q := c.Request.URL.Query()
	q.Add("provider", provider)
	c.Request.URL.RawQuery = q.Encode()
	if gothUser, err := gothic.CompleteUserAuth(c.Writer, c.Request); err == nil {
		log.Printf("Utilisateur authentifié: %s", gothUser.Email)
		c.JSON(http.StatusOK, gin.H{"user": gothUser})
	} else {
		log.Printf("Début de l'authentification avec le provider %s", provider)
		gothic.BeginAuthHandler(c.Writer, c.Request)
	}
}

func GetAuthCallback(c *gin.Context) {
	provider := c.Param("provider")
	log.Printf("callback authentification avec le provider %s", provider)
	q := c.Request.URL.Query()
	q.Add("provider", provider)
	c.Request.URL.RawQuery = q.Encode()
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		log.Printf("Erreur lors de l'authentification: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Utilisateur authentifié: %s", user.Email)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func Logout(c *gin.Context) {
	if err := gothic.Logout(c.Writer, c.Request); err != nil {
		log.Printf("Erreur lors de la déconnexion: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Writer.Header().Set("Location", "/")
	c.Writer.WriteHeader(http.StatusTemporaryRedirect)
}
