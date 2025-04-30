package controllers

import (
	"llio-api/models/DTOs"
	"llio-api/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var projectSTR = "projet"

func CreatedProject(c *gin.Context) {
	var projetDTO DTOs.ProjectDTO

	msgErrsJson := services.VerifyJSON(c, &projetDTO)
	if len(msgErrsJson) > 0 {
		log.Printf("Une ou plusieurs erreurs de format JSON sont survenues:%v", msgErrsJson)
		c.JSON(http.StatusBadRequest, gin.H{"errors": msgErrsJson})
		return
	}

	msgErrsMore := services.VerifyProjectJSON(&projetDTO)
	if len(msgErrsMore) > 0 {
		log.Printf("Une ou plusieurs erreurs de verification du format du projet sont survenues:%v", msgErrsMore)
		c.JSON(http.StatusBadRequest, gin.H{"errors": msgErrsMore})
		return
	}

	_, err := services.GetUserById(strconv.Itoa(projetDTO.ManagerId))
	if err != nil {
		handleError(c, err, userSTR)
		return
	}

	projectAdded, err := services.CreateProject(&projetDTO)
	if err != nil {
		handleError(c, err, projectSTR)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"response": "Le projet a bien été ajouté à la base de données",
		"project":  projectAdded,
	})
}

func GetProjectById(c *gin.Context) {
	id := c.Param("id")

	project, err := services.GetProjectById(id)
	if err != nil {
		handleError(c, err, projectSTR)
		return
	}

	c.JSON(http.StatusOK, gin.H{"project": project})
}

func GetProjects(c *gin.Context) {
	projects, err := services.GetProjects()
	if err != nil {
		handleError(c, err, projectSTR)
		return
	}

	c.JSON(http.StatusOK, gin.H{"projets": projects})
}

func UpdateProject(c *gin.Context) {
	var projectToUpdate DTOs.ProjectDTO

	msgErrsJson := services.VerifyJSON(c, &projectToUpdate)
	if len(msgErrsJson) > 0 {
		log.Printf("Une ou plusieurs erreurs de format JSON sont survenues:%v", msgErrsJson)
		c.JSON(http.StatusBadRequest, gin.H{"errors": msgErrsJson})
		return
	}

	msgErrsMore := services.VerifyProjectJSON(&projectToUpdate)
	if len(msgErrsMore) > 0 {
		log.Printf("Une ou plusieurs erreurs de verification du format du projet sont survenues:%v", msgErrsMore)
		c.JSON(http.StatusBadRequest, gin.H{"errors": msgErrsMore})
		return
	}

	id := strconv.Itoa(projectToUpdate.Id)
	_, err := services.GetProjectById(id)
	if err != nil {
		handleError(c, err, projectSTR)
		return
	}

	_, err = services.GetUserById(strconv.Itoa(projectToUpdate.ManagerId))
	if err != nil {
		handleError(c, err, userSTR)
		return
	}

	updatedProjectDTO, err := services.UpdateProject(&projectToUpdate)
	if err != nil {
		handleError(c, err, projectSTR)
		return
	}

	c.JSON(http.StatusOK, gin.H{"updatedProject": updatedProjectDTO})
}
