package middleware

import (
	"llio-api/models/DTOs"
	"llio-api/models/enums"
	"llio-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoleValidationMiddleware(requiredRole enums.UserRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Cookie("access_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token in the cookie"})
			c.Abort()
			return
		}

		user, err := services.ValidateTokenAndExtractUser(accessToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token requis"})
			c.Abort()
			return
		}

		var shouldReturn bool
		user, shouldReturn = setUserToContextAndValidate(c, user)
		if shouldReturn {
			return
		}

		if user.Role < requiredRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "Accès refusé"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func setUserToContextAndValidate(c *gin.Context, user *DTOs.UserDTO) (*DTOs.UserDTO, bool) {
	c.Set("current_user", user)

	currentUser, exists := c.Get("current_user")
	if !exists {
		c.JSON(http.StatusForbidden, gin.H{"error": "Échec lors de la sauvegarde de l'utilisateur dans la session"})
		c.Abort()
		return nil, true
	}

	user, ok := currentUser.(*DTOs.UserDTO)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur interne du serveur"})
		c.Abort()
		return nil, true
	}
	return user, false
}
