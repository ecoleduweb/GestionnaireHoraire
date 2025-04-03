package controllers

import (
	"llio-api/useful"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func GetAuthCallback(c *gin.Context) {
	useful.LoadEnv()
	frontendURL := os.Getenv("FRONTEND_URL")
	useful.SetupAuthProvider(c)
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		log.Printf("Erreur lors de l'authentification: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Utilisateur authentifié")
	log.Printf(user.AccessToken)
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "access_token",
		Value:    user.AccessToken,
		Path:     "/",
		HttpOnly: true,                 // Keep this true for security if it's an auth token
		Secure:   false,                // Set to true when using HTTPS
		Domain:   "",                   // Empty domain works better for localhost
		SameSite: http.SameSiteLaxMode, // SameSiteLaxMode is most secure, but requires browser support
	})
	http.Redirect(c.Writer, c.Request, frontendURL+"/calendar", http.StatusFound)
}

func Auth(c *gin.Context) {
	useful.SetupAuthProvider(c)
	if gothUser, err := gothic.CompleteUserAuth(c.Writer, c.Request); err == nil {
		log.Printf("Utilisateur déja authentifié")
		c.JSON(http.StatusOK, gin.H{"access_token": gothUser.AccessToken})
		return
	} else {
		log.Printf("Début de l'authentification")
		gothic.BeginAuthHandler(c.Writer, c.Request)
	}
}

func Logout(c *gin.Context) {
	useful.LoadEnv()
	frontendURL := os.Getenv("FRONTEND_URL")
	if err := gothic.Logout(c.Writer, c.Request); err != nil {
		log.Printf("Erreur lors de la déconnexion: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	http.Redirect(c.Writer, c.Request, frontendURL, http.StatusFound)
}
