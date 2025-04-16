package services

import (
	"errors"
	"llio-api/models/DTOs"
	"llio-api/models/enums"
	"llio-api/useful"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func ValidateTokenAndExtractUser(accessToken string) (*DTOs.UserDTO, error) {
	useful.LoadEnv()

	// Validate the token and extract claims

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	// Extract claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("token invalide")
	}

	user := &DTOs.UserDTO{
		Id:        int(claims["user_id"].(float64)),
		Email:     claims["email"].(string),
		FirstName: claims["first_name"].(string),
		LastName:  claims["last_name"].(string),
		Role:      enums.UserRole(claims["role"].(float64)),
	}
	return user, nil
}

func CreateJWTToken(userId int, userEmail string, fisrtName string, lastName string, expiresAt time.Time, userRole enums.UserRole) (string, error) {
	// Define the claims for the token
	useful.LoadEnv()
	secretKey := os.Getenv("JWT_SECRET_KEY")
	claims := jwt.MapClaims{
		"user_id":    userId,
		"email":      userEmail,
		"first_name": fisrtName,
		"last_name":  lastName,
		"exp":        expiresAt.Unix(),
		"role":       userRole, // TODO aller chercher le role dans la base de données et le passer en paramètre ici
	}

	// Create a new token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
