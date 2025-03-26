package routes

import (
	"github.com/gin-gonic/gin"
	//Importation des controleurs
	"llio-api/controllers"
)

func RegisterRoutes(r *gin.Engine) {
	/*------------------- Activities -------------------*/
	activityGroup := r.Group("/activity")
	{
		activityGroup.POST("", controllers.CreateActivity)
		activityGroup.GET("/:id", controllers.GetActivityById)
		activityGroup.PUT("", controllers.UpdateActivity)
		activityGroup.DELETE("/:id", controllers.DeleteActivity)
	}

	activitiesGroup := r.Group("/activities")
	{
		activitiesGroup.GET("", controllers.GetAllActivities)
	}

	/*------------------- Categories -------------------*/
	categoryGroup := r.Group("/category")
	{
		categoryGroup.POST("", controllers.CreateCategory)
		categoryGroup.GET("/:id", controllers.GetCategoryById)
	}

	categoriesGroup := r.Group("/categories")
	{
		categoriesGroup.GET("", controllers.GetCategories)
	}

}

// Routes pour l'authentification
func AuthRoutes(r *gin.Engine) {
	r.GET("/auth/:provider/callback", controllers.GetAuthCallback)
	r.GET("/logout/:provider", controllers.Logout)
	r.GET("/auth/:provider", controllers.Auth)
}
