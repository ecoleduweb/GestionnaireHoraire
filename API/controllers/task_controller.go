package controllers

import (
	"errors"
	"llio-api/models/DTOs"
	"llio-api/services"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateTask(c *gin.Context) {

	var taskDTO DTOs.TaskDTO

	//Validation des données entrantes
	messageErrsJSON := services.VerifyJSON(c, &taskDTO)
	if len(messageErrsJSON) > 0 {
		log.Printf("Une ou plusieurs erreurs de format JSON sont survenues:%v", err)
		c.JSON(http.StatusBadRequest, gin.H{"errors": messageErrsJSON})
		return
	}

	messageErrs := services.VerifyCreateTaskJSON(&taskDTO)
	if len(messageErrs) > 0 {
		log.Printf("Une ou plusieurs erreurs de verification du format de la tâche sont survenues:%v", messageErrs)
		c.JSON(http.StatusBadRequest, gin.H{"errors": messageErrs})
		return
	}

	task, err := services.CreateTask(&taskDTO)
	if err != nil {
		log.Printf("Erreur critique du server:%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"reponse": "La tâche a bien été ajoutée à la base de données.",
		"task":    task,
	})
}

func GetAllTasks(c *gin.Context) {

	tasks, err := services.GetAllTasks()
	if err != nil {
		log.Printf("Impossible de récupérer les tâches:%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de récupérer les tâches."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func GetTaskById(c *gin.Context) {
	// Récupérer l'id de la tâche
	id := c.Param("id")
	task, err := services.GetTaskById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("La tâche est introuvable:%v", err)
			c.JSON(http.StatusNotFound, gin.H{"error": "La tâche est introuvable."})
			return
		}
		log.Printf("Impossible de récupérer les tâches:%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de récupérer la tâche."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"task": task})
}
