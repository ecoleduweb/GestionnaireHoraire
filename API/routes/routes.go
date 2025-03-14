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

	activityGroup := r.Group("/activity")
	{
		activityGroup.POST("", controllers.CreateActivity)
		activityGroup.Get("/all", controllers.GetAllActivities)
		activityGroup.Get("/:id", controllers.GetActivityById)
		activityGroup.PUT("", controllers.UpdateActivity)
	}

}

// Routes pour l'authentification
func AuthRoutes(r *gin.Engine) {
	r.GET("/auth/:provider/callback", controllers.GetAuthCallback)
	r.GET("/logout/:provider", controllers.Logout)
	r.GET("/auth/:provider", controllers.Auth)
}
