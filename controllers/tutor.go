package controllers

import (
	"adopet/database"
	"adopet/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllTutores(c *gin.Context) {
	var tutores []models.Tutor
	database.DB.Find(&tutores)

	c.JSON(http.StatusOK, tutores)

}

func CreateTutor(c *gin.Context) {
	var tutor models.Tutor
	if err := c.ShouldBindJSON(&tutor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&tutor)
	c.JSON(http.StatusOK, tutor)
}

func AlterTutor(c *gin.Context) {
	var tutor models.Tutor
	id := c.Params.ByName("id")

	database.DB.First(&tutor, id)

	if err := c.ShouldBindJSON(&tutor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Model(&tutor).UpdateColumns(tutor)
	c.JSON(http.StatusOK, tutor)
}

func GetIDTutor(c *gin.Context) {
	var tutor models.Tutor
	id := c.Params.ByName("id")
	database.DB.First(&tutor, id)

	if tutor.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Tutor não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, tutor)
}

func DeleteTutor(c *gin.Context) {
	var tutor models.Tutor
	id := c.Params.ByName("id")

	database.DB.First(&tutor, id)

	if tutor.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Tutor não encontrado",
		})
		return
	}

	database.DB.Delete(&tutor, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Tutor deletado com sucesso!",
	})
}
