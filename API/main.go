package main

import (
	"llio-api/database"
	"llio-api/routes"

	"github.com/gin-gonic/gin"
)

// Chercher à améliorer et sécuriser
// Version fonctionnel mais inspirer de ChatGPT sans vérification de sécurité et efficacité
func main() {
	r := gin.Default()
	database.Connect()
	routes.RegisterRoutes(r)
	r.Run(":8080")
}
