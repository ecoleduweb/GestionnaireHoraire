package controllers

import (
	"errors"
	"llio-api/models/DTOs"
	"llio-api/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
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

	// Verifier si l'utilisateur existe et est authentifié (middleware)

	// Vérifier que le projet existe et ne soit pas cloturé

	categoryAdded, err := services.CreateCategory(&categoryDTO)
	if err != nil {
		// Vérification spécifique pour MySQL
		if is, ok := err.(*mysql.MySQLError); ok {
			// Code d'erreur MySQL pour entrée dupliquée
			if is.Number == 1062 {
				c.JSON(http.StatusConflict, gin.H{
					"reponse": "Une catégorie du même nom existe déjà pour ce projet.",
				})
				return
			}
		}

		log.Printf("Erreur critique du server:%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
		log.Printf("Impossible de récupérer les catégories:%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de récupérer les catégories."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"categories": categories})
}

func GetCategoryById(c *gin.Context) {
	id := c.Param("id")

	category, err := services.GetCategoryById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("La catégorie est introuvable:%v", err)
			c.JSON(http.StatusNotFound, gin.H{"error": "La catégorie est introuvable."})
			return
		}
		log.Printf("Impossible de récupérer la catégorie:%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de récupérer la catégorie."})
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("La catégorie est introuvable:%v", err)
			c.JSON(http.StatusNotFound, gin.H{"error": "La catégorie est introuvable."})
			return
		}
		log.Printf("Impossible de récupérer la catégorie:%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de récupérer la catégorie."})
		return
	}

	updatedCategoryDTO, err := services.UpdateCategory(&updateCategoryDTO)
	if err != nil {
		log.Printf("La catégorie n'a pas été modifiée : %v", err)
		c.JSON(http.StatusNotModified, gin.H{"error": "La catégorie n'a pas été modifiée."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"updatedCategory": updatedCategoryDTO})
}
