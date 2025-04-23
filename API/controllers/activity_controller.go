package controllers

import (
	"llio-api/models/DTOs"
	"llio-api/services"
	"log"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

var activiteSTR = "activité"

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
		log.Printf("Une ou plusieurs erreurs de verification du format de l'activité sont survenues:%v", messageErrs)
		c.JSON(http.StatusBadRequest, gin.H{"errors": messageErrs})
		return
	}

	activityAded, err := services.CreateActivity(&activityDTO)
	if err != nil {
		handleError(c, err, activiteSTR)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"reponse":  "L'activité a bien été ajoutée à la base de données.",
		"activity": activityAded,
	})
}

func GetUsersActivities(c *gin.Context) {

	currentUser, exists := c.Get("current_user")
	if !exists {
		c.JSON(http.StatusForbidden, gin.H{"error": "Échec lors de la sauvegarde de l'utilisateur dans la session"})
		return
	}

	user, ok := currentUser.(*DTOs.UserDTO)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur interne du serveur"})
		return
	}
	activities, err := services.GetUsersActivities(user.Id)
	if err != nil {
		handleError(c, err, activiteSTR)
		return
	}

	// Retourne une liste vide plutôt que null pour être compatible avec les vérifiacations du frontend
	if activities == nil {
		// Retourner une liste vide au lieu de null
		c.JSON(http.StatusOK, gin.H{"activities": []DTOs.ActivityDTO{}})
		return
	}

	c.JSON(http.StatusOK, gin.H{"activities": activities})
}

func GetActivityById(c *gin.Context) {
	// Récupérer l'id de l'activité
	id := c.Param("id")
	activity, err := services.GetActivityById(id)
	if err != nil {
		handleError(c, err, activiteSTR)
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
		handleError(c, err, activiteSTR)
		return
	}

	updatedActivityDTO, err := services.UpdateActivity(&updateActivityDTO)
	if err != nil {
		handleError(c, err, activiteSTR)
		return
	}

	c.JSON(http.StatusOK, gin.H{"updated_activity": updatedActivityDTO})

}

func DeleteActivity(c *gin.Context) {

	id := c.Param("id")
	activity, err := services.GetActivityById(id)
	if err != nil {
		handleError(c, err, activiteSTR)
		return
	}
	if activity == nil {
		log.Printf("L'activité à supprimer n'existe pas. : %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "L'activité à supprimer n'existe pas."})
		return
	}

	err = services.DeleteActivity(id)
	if err != nil {
		handleError(c, err, activiteSTR)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "La suppression de l'activité est un succès."})
}
