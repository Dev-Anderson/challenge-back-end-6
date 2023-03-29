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
