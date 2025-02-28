package controllers

import (
	"llio-api/models/DTOs"
	"llio-api/services"
	"log"
	"strconv"

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

	taskDTOAded, err := services.CreateTask(&taskDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"reponse": "La tâche a bien été ajoutée à la base de données.",
		"task":    taskDTOAded,
	})
}

// GetAllTasks récupère toutes les tâches
func GetAllTasks(c *gin.Context) {

	tasks, err := services.GetAllTasks()
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
	task, err := services.GetTaskById(id)
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

// UpdateTask permet de mettre à jour une tâche
func UpdateTask(c *gin.Context) {
	var updateTaskDTO DTOs.TaskResponse

	//Validation des données entrantes
	messageErrsJSON := services.VerifyJSON(c, &updateTaskDTO)
	if len(messageErrsJSON) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": messageErrsJSON})
		return
	}

	id := strconv.Itoa(updateTaskDTO.Id)
	task, err := services.GetTaskByIdService(id)
	if err != nil {
		log.Printf("Impossible de récupérer la tâche:%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de récupérer la tâche."})
		return
	}
	if task == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "La tâche est introuvable."})
		return
	}

	err = services.UpdateTaskService(&updateTaskDTO)
	if err != nil {
		log.Printf("La tâche n'a pas été modifiée : %v", err)
		c.JSON(http.StatusNotModified, gin.H{"error": "La tâche n'a pas été modifiée."})
		return
	}

	updatedTaskDTO, err := services.GetTaskByIdService(id)
	if err != nil {
		log.Printf("Impossible de récupérer la tâche:%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de récupérer la tâche."})
		return
	}
	if updatedTaskDTO == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "La tâche est introuvable."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"updated_task": updatedTaskDTO})

}
