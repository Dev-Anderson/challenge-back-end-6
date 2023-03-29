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
