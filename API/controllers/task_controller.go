package controllers

import (
	"llio-api/models/DTOs"
	"llio-api/services"
	"log"

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

// GetAllTasks récupère toutes les tâches
func GetAllTasks(c *gin.Context) {

	tasks, err := services.GetAllTasksService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de récupérer les tâches."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

// GetTaskById récupère une tâche par son id
func GetTaskById(c *gin.Context) {
	// Récupérer l'id de la tâche
	id := c.Param("id")
	task, err := services.GetTaskByIdService(id)
	if err != nil {
		log.Printf("Impossible de récupérer les tâches:%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de récupérer la tâche."})
		return
	}
	if task == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "La tâche est introuvable."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"task": task})
}
