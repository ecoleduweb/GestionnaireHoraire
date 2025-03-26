package controllers

import (
	"llio-api/models/DTOs"
	"llio-api/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {

	var categoryDTO DTOs.CategoryDTO

	messageErrsJSON := services.VerifyJSON(c, &categoryDTO)
	if len(messageErrsJSON) > 0 {
		log.Printf("Une ou plusieurs erreurs de format JSON sont survenues:%v", messageErrsJSON)
		c.JSON(http.StatusBadRequest, gin.H{"errors": messageErrsJSON})
		return
	}

	messageErrs := services.VerifyCreateCategoryJSON(&categoryDTO)
	if len(messageErrs) > 0 {
		log.Printf("Une ou plusieurs erreurs de verification du format de l'activité sont survenues:%v", messageErrs)
		c.JSON(http.StatusBadRequest, gin.H{"errors": messageErrs})
		return
	}

	categoryAdded, err := services.CreateCategory(&categoryDTO)
	if err != nil {
		log.Printf("Erreur critique du server:%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"reponse":  "La catégorie a bien été ajoutée à la base de données.",
		"activity": categoryAdded,
	})
}
