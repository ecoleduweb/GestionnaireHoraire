package middleware

import (
	"errors"
	"llio-api/models/DAOs"
	"llio-api/models/enums"
	"llio-api/useful"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func validateTokenAndSetUserToContext(c *gin.Context) error {
	useful.LoadEnv()

	accessToken, err := c.Cookie("access_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token requis"})
		c.Abort()
		return err
	}
	// Validate the token and extract claims

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token invalide"})
		c.Abort()
		return err
	}

	// Extract claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token invalide"})
		c.Abort()
		return errors.New("token invalide")
	}

	user := DAOs.User{
		Email:     claims["email"].(string),
		FirstName: claims["first_name"].(string),
		LastName:  claims["last_name"].(string),
		Role:      enums.UserRole(int(claims["role"].(float64))),
	}
	c.Set("current_user", user)
	return nil
}

func RoleValidationMiddleware(requiredRole enums.UserRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := validateTokenAndSetUserToContext(c); err != nil {
			return
		}

		// Retrieve the current user from the context
		currentUser, exists := c.Get("current_user")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "Utilisateur non authentifié"})
			c.Abort()
			return
		}

		// Cast the current user to the expected type
		user, ok := currentUser.(DAOs.User)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur interne du serveur"})
			c.Abort()
			return
		}

		// Check if the user's role matches the required role
		if user.Role != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "Accès refusé"})
			c.Abort()
			return
		}

		// If the role is valid, continue the request
		c.Next()
	}
}
