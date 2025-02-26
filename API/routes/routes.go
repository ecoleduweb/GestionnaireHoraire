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
		// Route pour la création d'une tâche
		taskGroup.POST("", controllers.CreateTask)

<<<<<<< Updated upstream
		//Route pour la récupération de toutes les tâches
		taskGroup.GET("", controllers.GetAllTasks)

		//Route pour la récupération d'une tâche par son id
		taskGroup.GET("/:id", controllers.GetTaskById)

		//Route pour mettre à jour une tâche
		taskGroup.PUT("", controllers.UpdateTask)

=======
		// Route pour mettre à jour une tâche
		taskGroup.PUT("/:id", controllers.UpdateTask)
>>>>>>> Stashed changes
	}

}
