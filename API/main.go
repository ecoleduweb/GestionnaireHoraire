package main

import (
	"llio-api/auth"
	"llio-api/database"
	"llio-api/handlers"
	"llio-api/routes"
	"llio-api/useful"
	"os"

	"llio-api/auth"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Chercher à améliorer et sécuriser
// Version fonctionnel mais inspirer de ChatGPT sans vérification de sécurité et efficacité
func main() {
	auth.NewAuth()

	useful.LoadEnv()
	if os.Getenv("ENV") == "PROD" {
		gin.SetMode(gin.ReleaseMode) // Désactive les logs de débogage en production
	}

	r := gin.Default()

	frontendAddress := os.Getenv("FRONT_ADDRESS")
	if frontendAddress == "" {
		frontendAddress = "http://localhost:5173" // Port par défaut
	}

	// Configuration CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{frontendAddress},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 3600, // Durée de mise en cache de la requête preflight (en secondes)
	}))

	routes.RegisterRoutes(r)
	handlers.ApiStatus(r)
	routes.AuthRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Port par défaut
	}

	database.Connect()

	r.Run(":" + port)
}
