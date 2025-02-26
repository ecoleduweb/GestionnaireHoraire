package controllers

import (
	"llio-api/models/DTOs"
	"llio-api/services"
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

<<<<<<< Updated upstream
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	if task == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Impossible de récupérer la tâche."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"task": task})
}

// UpdateTask permet de mettre à jour une tâche
func UpdateTask(c *gin.Context) {
	var updateTaskDTO DTOs.TaskResponse

	//Validation des données entrantes
	messageErrsJSON := services.VerifyJSON(c, &updateTaskDTO)
=======
func UpdateTask(c *gin.Context) {
	var taskDTO DTOs.TaskDTO

	//Validation des données entrantes
	messageErrsJSON := services.VerifyJSON(c, &taskDTO)
>>>>>>> Stashed changes
	if len(messageErrsJSON) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": messageErrsJSON})
		return
	}

<<<<<<< Updated upstream
	id := strconv.Itoa(updateTaskDTO.Id)
	task, err := services.GetTaskByIdService(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	if task == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Impossible de récupérer la tâche."})
		return
	}

	err = services.UpdateTaskService(&updateTaskDTO)
	if err != nil {
		c.JSON(http.StatusNotModified, gin.H{"error": err})
		return
	}

	updatedTaskDTO, err := services.GetTaskByIdService(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	if updatedTaskDTO == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Impossible de récupérer la tâche."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"updated_task": updatedTaskDTO})

=======
	messageErrs := services.VerifyCreateTaskJSON(&taskDTO)
	if len(messageErrs) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": messageErrs})
		return
	}

	taskUpdated, err := services.UpdateTaskService(&taskDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"reponse": "La tâche a bien été mise à jour dans la base de données.",
		"task":    taskUpdated,
	})
>>>>>>> Stashed changes
}
