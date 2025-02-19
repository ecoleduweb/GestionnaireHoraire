package routes

import (
	"github.com/gin-gonic/gin"
	//Importation des controleurs
	"llio-api/controllers"
)

func RegisterRoutes(r *gin.Engine) {
	/*
		Route de création de tâche

		Pour l'instant aucune vérification de l'idetification d'un usager n'est faite!
		( 2025-01-29/Quentin ): Il faudra faire appel à un middleware pour valider le token d'authentification
	*/

	r.POST("/create-user", controllers.CreateTask)

	// Groupe de routes pour les tâches(event/task)
	taskGroup := r.Group("/tasks")
	{
		taskGroup.POST("", controllers.CreateTask)
	}

}

func ApiStatus(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "pong",
		})
	})

	r.GET("/health/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "Healthy",
		})
	})
}
