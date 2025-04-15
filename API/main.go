package main

import (
	"io"
	"llio-api/auth"
	"llio-api/cmd"
	"llio-api/database"
	"llio-api/handlers"
	"llio-api/middleware"
	"llio-api/routes"
	"llio-api/useful"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Chercher à améliorer et sécuriser
// Version fonctionnel mais inspirer de ChatGPT sans vérification de sécurité et efficacité
func main() {
	// permet de passer des commandes de migraion de la base de données
	// ex : go run main.go migrate create initial_migration
	if len(os.Args) > 1 {
		cmd.Execute()
		return
	}

	logFile, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	defer logFile.Close()

	//Pour copier les logs dans le fichier en plus de la console
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter)
	log.SetFlags(log.Ldate | log.Ltime)

	gin.DefaultWriter = multiWriter
	gin.DefaultErrorWriter = multiWriter

	auth.AuthWithAzure()

	useful.LoadEnv()
	if os.Getenv("ENV") == "PROD" {
		gin.SetMode(gin.ReleaseMode) // Désactive les logs de débogage en production
	}

	r := gin.Default()

	frontendAddress := os.Getenv("FRONT_ADDRESS")
	if frontendAddress == "" {
		frontendAddress = "http://localhost:5173" // Port par défaut
	}

	if os.Getenv("ENABLE_TRACING") == "true" {
		shutdown := useful.InitTracer()
		defer shutdown()
	}

	r.Use(middleware.Telemetry())

	// Configuration CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{frontendAddress},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept", "traceparent", "tracestate"},
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
