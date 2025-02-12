package controllers

import (
	"llio-api/models/DTOs"
	"llio-api/services"

	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {

	var taskDTO DTOs.TaskDTO

	//Validation des données entrantes
	if err := c.ShouldBindJSON(&taskDTO); err != nil {
		println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Une ou plusieurs données sont invalides ou manquantes."})
		return
	}

	if err := services.CreateTaskService(&taskDTO); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"reponse": "La tâche a bien été ajoutée à la base de données."})
}
