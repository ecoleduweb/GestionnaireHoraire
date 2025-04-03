package routes

import (
	"github.com/gin-gonic/gin"
	//Importation des controleurs
	"llio-api/controllers"
	"llio-api/middleware"
	"llio-api/models/enums"
)

// Importation du middleware
func RegisterRoutes(r *gin.Engine) {
	/*------------------- Activities -------------------*/
	activityGroup := r.Group("/activity", middleware.RoleValidationMiddleware((enums.Employee)))
	{
		activityGroup.POST("", controllers.CreateActivity)
		activityGroup.GET("/:id", controllers.GetActivityById)
		activityGroup.PUT("", controllers.UpdateActivity)
		activityGroup.DELETE("/:id", controllers.DeleteActivity)
	}

	activitiesGroup := r.Group("/activities", middleware.RoleValidationMiddleware(enums.Employee))
	{
		activitiesGroup.GET("", controllers.GetAllActivities)
	}

	/*------------------- Categories -------------------*/
	categoryGroup := r.Group("/category", middleware.RoleValidationMiddleware(enums.Employee))
	{
		categoryGroup.POST("", controllers.CreateCategory)
		categoryGroup.GET("/:id", controllers.GetCategoryById)
		categoryGroup.PUT("", controllers.UpdateCategory)
	}

	categoriesGroup := r.Group("/categories", middleware.RoleValidationMiddleware(enums.Employee))
	{
		categoriesGroup.GET("", controllers.GetCategories)
	}

	/*------------------- Projets -------------------*/
	projectGroup := r.Group("/project")
	{
		projectGroup.POST("", controllers.CreatedProject)
		projectGroup.GET("/:id", controllers.GetProjectById)
		projectGroup.PUT("", controllers.UpdateProject)
	}

	projectsGroup := r.Group("/projects")
	{
		projectsGroup.GET("", controllers.GetProjects)
	}

}

// Routes pour l'authentification
func AuthRoutes(r *gin.Engine) {
	r.GET("/auth/:provider/callback", controllers.GetAuthCallback)
	r.GET("/logout/:provider", controllers.Logout)
	r.GET("/auth/:provider", controllers.Auth)
}
