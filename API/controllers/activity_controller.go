package controllers

import (
	"errors"
	"llio-api/models/DTOs"
	"llio-api/services"
	"log"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateActivity(c *gin.Context) {

	var activityDTO DTOs.ActivityDTO

	//Validation des données entrantes
	messageErrsJSON := services.VerifyJSON(c, &activityDTO)
	if len(messageErrsJSON) > 0 {
		log.Printf("Une ou plusieurs erreurs de format JSON sont survenues:%v", messageErrsJSON)
		c.JSON(http.StatusBadRequest, gin.H{"errors": messageErrsJSON})
		return
	}

	messageErrs := services.VerifyCreateActivityJSON(&activityDTO)
	if len(messageErrs) > 0 {
		log.Printf("Une ou plusieurs erreurs de verification du format de la tâche sont survenues:%v", messageErrs)
		c.JSON(http.StatusBadRequest, gin.H{"errors": messageErrs})
		return
	}

	activityAded, err := services.CreateActivityService(&activityDTO)
	if err != nil {
		log.Printf("Erreur critique du server:%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"reponse":  "L'activité a bien été ajoutée à la base de données.",
		"activity": activityAded,
	})
}

func GetAllActivities(c *gin.Context) {

	activities, err := services.GetAllActivities()
	if err != nil {
		log.Printf("Impossible de récupérer les tâches:%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de récupérer les tâches."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"activities": activities})
}

func GetActivityById(c *gin.Context) {
	// Récupérer l'id de la tâche
	id := c.Param("id")
	activity, err := services.GetActivityById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("La tâche est introuvable:%v", err)
			c.JSON(http.StatusNotFound, gin.H{"error": "La tâche est introuvable."})
			return
		}
		log.Printf("Impossible de récupérer la tâche:%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de récupérer la tâche."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"activity": activity})
}

func UpdateActivity(c *gin.Context) {
	var updateActivityDTO DTOs.ActivityDTO

	//Validation des données entrantes
	messageErrsJSON := services.VerifyJSON(c, &updateActivityDTO)
	if len(messageErrsJSON) > 0 {
		log.Printf("Une ou plusieurs erreurs de format JSON sont survenues:%v", messageErrsJSON)
		c.JSON(http.StatusBadRequest, gin.H{"errors": messageErrsJSON})
		return
	}

	id := strconv.Itoa(updateActivityDTO.Id)
	_, err := services.GetActivityById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("La tâche à mettre à jour est introuvable:%v", err)
			c.JSON(http.StatusNotFound, gin.H{"error": "La tâche est introuvable."})
			return
		}
		log.Printf("Impossible de récupérer la tâche à mettre à jour:%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de récupérer la tâche à mettre à jour."})
		return
	}

	updatedActivityDTO, err := services.UpdateActivity(&updateActivityDTO)
	if err != nil {
		log.Printf("La tâche n'a pas été modifiée : %v", err)
		c.JSON(http.StatusNotModified, gin.H{"error": "La tâche n'a pas été modifiée."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"updated_activity": updatedActivityDTO})

}
