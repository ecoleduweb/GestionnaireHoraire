package controllers

import (
	"llio-api/models/DTOs"
	"llio-api/services"

	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateActivity(c *gin.Context) {

	var activityDTO DTOs.ActivityDTO

	//Validation des données entrantes
	messageErrsJSON := services.VerifyJSON(c, &activityDTO)
	if len(messageErrsJSON) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": messageErrsJSON})
		return
	}

	messageErrs := services.VerifyCreateActivityJSON(&activityDTO)
	if len(messageErrs) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": messageErrs})
		return
	}

	activityAded, err := services.CreateActivityService(&activityDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"reponse":  "La tâche a bien été ajoutée à la base de données.",
		"activity": activityAded,
	})
}
