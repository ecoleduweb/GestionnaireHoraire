package controllers

import (
	"llio-api/models/DTOs"
	"llio-api/services"

	"net/http"

	"github.com/gin-gonic/gin"
)

// Fonction controlleur qui permet d'ajouter une tâche dans la base de données.
func CreateTask(c *gin.Context) {
	//Récupération de la tâche dans la requête

	var reqBody DTOs.TaskDTO

	//Validation des données entrantes
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Une ou plusieurs données sont invalides ou manquantes."})
		return
	}

	//Appeler le service pour créer une tâche
	task, err := services.CreateTaskService(&reqBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//Retourne la réponse au frontend
	c.JSON(http.StatusOK, gin.H{"reponse": "La tâche a bien été ajoutée à la base de données.", "task": task})
}
