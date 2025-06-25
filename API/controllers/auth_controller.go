package controllers

import (
	"llio-api/models/DTOs"
	"llio-api/models/enums"
	"llio-api/services"
	"llio-api/useful"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func GetAuthCallback(c *gin.Context) {
	useful.LoadEnv()
	frontendURL := os.Getenv("FRONTEND_URL")
	useful.SetupAuthProvider(c)
	userAzure, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		log.Printf("Erreur lors de l'authentification: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var dbUser DTOs.UserDTO
	dbUser.FirstName = userAzure.FirstName
	dbUser.LastName = userAzure.LastName
	dbUser.Email = userAzure.Email
	dbUser.Role = enums.Employee

	userInDb, err := services.FirstOrCreateUser(&dbUser)
	if err != nil {
		log.Printf("Impossible d'interagir avec l'utilisateur dans la base de données: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible d'ajouter l'utilisateur à la base de données."})
		return
	}

	accessToken, err := services.CreateJWTToken(userInDb.Id, userAzure.Email, userAzure.FirstName, userAzure.LastName, userAzure.ExpiresAt, userInDb.Role)
	if err != nil {
		log.Printf("Erreur lors de l'authentification: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	isRunningSecure := useful.IsRunningSecure()

	// cookie pour le graphapi!
	// Utilisé quand on va importer les valeurs du calendrier de l'utilisateur
	// http.SetCookie(c.Writer, &http.Cookie{
	// 	Name:     "graphapi_access_token",
	// 	Value:    user.AccessToken,
	// 	Path:     "/",
	// 	HttpOnly: true,
	// 	Secure:   isRunningSecure,
	// 	SameSite: http.SameSiteStrictMode,
	// })

	// cookie pour l'authentification de l'utilisateur
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   isRunningSecure,
		SameSite: http.SameSiteStrictMode,
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
	if err := gothic.Logout(c.Writer, c.Request); err != nil {
		log.Printf("Erreur lors de la déconnexion: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "access_token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   useful.IsRunningSecure(),
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
	})

	c.JSON(http.StatusOK, gin.H{"message": "Déconnexion réussie"})
}
