package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Récupérer le token du header Authorization
		// Récupérer le token du cookie access_token si le header Authorization est vide
		//TODO passer par le header Authorization i guess.... :(
		for _, cookie := range c.Request.Cookies() {
			fmt.Printf("Cookie found: %s = %s\n", cookie.Name, cookie.Value)
		}

		accessToken, err := c.Cookie("access_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token requis"})
			c.Abort()
			return
		}
		// Vérifier le format du token
		// Valider le token
		_, err = ValidateMicrosoftToken(accessToken, c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token invalide"})
			c.Abort()
			return
		}

		// Si le token est valide, continuer la requête
		c.Next()
	}
}
