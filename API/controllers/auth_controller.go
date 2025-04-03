package controllers

import (
	"llio-api/auth"
	"llio-api/models/enums"
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

	// TODO : vérifier si l'utilisateur existe dans la base de données et aller chercher son role.
	// Si le user n'existe pas, le créer avec le role par défaut (user).

	accessToken, err := auth.CreateJWTToken(user.Email, user.FirstName, user.LastName, user.ExpiresAt, enums.Employee)
	if err != nil {
		log.Printf("Erreur lors de l'authentification: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// cookie pour le graphapi
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "graphapi_access_token",
		Value:    user.AccessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,                    // TODO env
		Domain:   "",                       // TODO env
		SameSite: http.SameSiteDefaultMode, // SameSiteLaxMode is most secure, but requires browser support
	})

	// cookie pour l'authentification de l'utilisateur
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,                    // TODO env
		Domain:   "",                       // TODO env
		SameSite: http.SameSiteDefaultMode, // SameSiteLaxMode is most secure, but requires browser support
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
