package controllers

import (
	"llio-api/database"
	"llio-api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Fonction controlleur qui permet d'ajouter une tâche dans la base de données.
func CreateTask(c *gin.Context) {
	//Récupération de la tâche dans la requête

	var reqBody struct {
		Name        string    `json:"name" binding:"required"`
		Description string    `json:"description" binding:"required"`
		State       string    `json:"state"`
		Billable    bool      `json:"billable"`
		StartDate   time.Time `json:"start_date" binding:"required"`
		EndDate     time.Time `json:"end_date" binding:"required"`
		//On va le trouver avec le token UserId	mais pour l'instant
		UserId     int `json:"user_id" binding:"required"`
		ProjectId  int `json:"project_id" binding:"required"`
		CategoryId int `json:"category_id" binding:"required"`
	}

	//Validation des données entrantes
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Une ou plusieurs données sont invalides ou manquantes."})
		return
	}

	//Création d'une tâche
	task := models.Task{
		Name:        reqBody.Name,
		Description: reqBody.Description,
		State:       reqBody.State,
		Billable:    reqBody.Billable,
		StartDate:   reqBody.StartDate,
		EndDate:     reqBody.EndDate,
		UserId:      reqBody.UserId,
		ProjectId:   reqBody.ProjectId,
		CategoryId:  reqBody.CategoryId,
	}

	//Vérification valeur par défault
	if task.State == "" {
		task.State = "à faire"
	}

	//Création de la tâches dans la base de données
	if err := database.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de créer la tâche dans la base de données."})
		return
	}

	//Retourne la réponse au frontend
	c.JSON(http.StatusOK, gin.H{"reponse": "La tâche à bien été ajouté à la base de données."})
}
