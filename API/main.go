package main

import (
	"llio-api/database"
	"llio-api/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Chercher à améliorer et sécuriser
// Version fonctionnel mais inspirer de ChatGPT sans vérification de sécurité et efficacité
func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("Aucun fichier .env trouvé, utilisation des variables d'environnement système")
	}
	if os.Getenv("ENV") == "PROD" {
		gin.SetMode(gin.ReleaseMode) // Désactive les logs de débogage en production
	}

	r := gin.Default()

	routes.RegisterRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Port par défaut
	}

	database.Connect()

	r.Run(":" + port)
}
