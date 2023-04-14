package controllers

import (
	"adopet/database"
	"adopet/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Criar um abrigo godoc
// @Summary 	Criar um novo perfil de Abrigo
// @Description	Criar um novo perfil de Abrigo
// @Tags 	Abrigo
// @Accept	json
// @Produce	json
// @Param 	Dados do Abrigo	body 	models.Abrigo	true 	"Informe os dados do Abrigo"
// @Success	200 {object} 	models.Abrigo
// @Router 	/abrigo		[post]
func CreateAbrigo(c *gin.Context) {
	var abrigo models.Abrigo
	if err := c.ShouldBindJSON(&abrigo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&abrigo)
	c.JSON(http.StatusOK, abrigo)
}

// Consulta todos os Abrigos godoc
// @Summary 	Consulta todos os Abrigos
// @Description 	Rota para consultar todos os Abrigos
// @Schemes
// @Tags 	Abrigo
// @Accet	json
// @Produce json
// @Success 	200 {object}	models.Abrigo
// @Router 	/abrigo		[get]
func GetAllAbrigo(c *gin.Context) {
	var abrigo []models.Abrigo
	database.DB.Find(&abrigo)
	c.JSON(http.StatusOK, abrigo)
}
