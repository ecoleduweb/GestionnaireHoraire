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
	messageErrsJSON := services.VerifyJSON(c, &taskDTO)
	if len(messageErrsJSON) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": messageErrsJSON})
		return
	}

	messageErrs := services.VerifyCreateTaskJSON(&taskDTO)
	if len(messageErrs) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": messageErrs})
		return
	}

	taskAded, err := services.CreateTaskService(&taskDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"reponse": "La tâche a bien été ajoutée à la base de données.",
		"task":    taskAded,
	})
}
