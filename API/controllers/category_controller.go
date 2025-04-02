package controllers

import (
	"llio-api/models/DTOs"
	"llio-api/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var categorieSTR = "catégorie"

func CreateCategory(c *gin.Context) {

	var categoryDTO DTOs.CategoryDTO

	messageErrsJSON := services.VerifyJSON(c, &categoryDTO)
	if len(messageErrsJSON) > 0 {
		log.Printf("Une ou plusieurs erreurs de format JSON sont survenues:%v", messageErrsJSON)
		c.JSON(http.StatusBadRequest, gin.H{"errors": messageErrsJSON})
		return
	}

	// Récupérer l'id du user depuis le token
	// A Modifier quand les user ok
	categoryDTO.UserId = 1

	messageErrs := services.VerifyCreateCategoryJSON(&categoryDTO)
	if len(messageErrs) > 0 {
		log.Printf("Une ou plusieurs erreurs de verification du format de la catégorie sont survenues:%v", messageErrs)
		c.JSON(http.StatusBadRequest, gin.H{"errors": messageErrs})
		return
	}

	categoryAdded, err := services.CreateCategory(&categoryDTO)
	if err != nil {
		handleDatabaseError(c, err, categorieSTR)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"reponse":  "La catégorie a bien été ajoutée à la base de données.",
		"activity": categoryAdded,
	})
}

func GetCategories(c *gin.Context) {
	categories, err := services.GetCategories()
	if err != nil {
		handleDatabaseError(c, err, categorieSTR)
		return
	}

	c.JSON(http.StatusOK, gin.H{"categories": categories})
}

func GetCategoryById(c *gin.Context) {
	id := c.Param("id")

	category, err := services.GetCategoryById(id)
	if err != nil {
		handleDatabaseError(c, err, categorieSTR)
		return
	}

	c.JSON(http.StatusOK, gin.H{"category": category})

}

func UpdateCategory(c *gin.Context) {
	var updateCategoryDTO DTOs.CategoryDTO

	//Validation des données entrantes
	messageErrsJSON := services.VerifyJSON(c, &updateCategoryDTO)
	if len(messageErrsJSON) > 0 {
		log.Printf("Une ou plusieurs erreurs de format JSON sont survenues:%v", messageErrsJSON)
		c.JSON(http.StatusBadRequest, gin.H{"errors": messageErrsJSON})
		return
	}

	id := strconv.Itoa(updateCategoryDTO.Id)
	_, err := services.GetCategoryById(id)
	if err != nil {
		handleDatabaseError(c, err, categorieSTR)
		return
	}

	updatedCategoryDTO, err := services.UpdateCategory(&updateCategoryDTO)
	if err != nil {
		handleDatabaseError(c, err, categorieSTR)
		return
	}

	c.JSON(http.StatusOK, gin.H{"updatedCategory": updatedCategoryDTO})
}
