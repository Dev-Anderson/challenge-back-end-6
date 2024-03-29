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

// Alterar Abrigo
// @Summary Alterar um abrigo existente.
// @Description Alterar um abrigo no banco de dados
// @ID alterAbrigo
// @Produce json
// @Tags Abrigo
// @Param id path int true "ID do abrigo a ser alterado"
// @Param abrigo body models.Abrigo true "Dados do abrigo a serem atualizado"
// @Success 200 {object} models.Abrigo
// @Failure 400 {object} error
// @Router /abrigo/{id} [patch]
func AlterAbrigo(c *gin.Context) {
	var abrigo models.Abrigo
	id := c.Params.ByName("id")

	database.DB.First(&abrigo, id)

	if err := c.ShouldBindJSON(&abrigo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Model(&abrigo).UpdateColumns(abrigo)
	c.JSON(http.StatusOK, abrigo)
}

// GetIDAbrigo recupera um abrigo por ID
// @Summary Recupera um abrigo por ID
// @Description Recupera um abrigo do banco de dados com base no ID
// @Tags Abrigo
// @Produce json
// @Param id path int true "ID do abrigo para consulta"
// @Success 200 {object} models.Abrigo
// @Failure 404 {object} error
// @Router /abrigo/{id} [get]
func GetIDAbrigo(c *gin.Context) {
	var abrigo models.Abrigo
	id := c.Params.ByName("id")
	database.DB.First(&abrigo, id)

	if abrigo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Abrigo nao encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, abrigo)
}

// DeleteAbrigo exclui um abrigo por DI
// @Summary Exclui um abrigo por ID
// @Description Exclui um abrigo do banco de dados com base no ID
// @Tags Abrigo
// @Produce json
// @Param id path int true "ID do abrigo para excluir"
// @Success 200 {object} string
// @Failure 404 {object} error
// @Router /abrigo/{id} [delete]
func DeleteAbrigo(c *gin.Context) {
	var abrigo models.Abrigo
	id := c.Params.ByName("id")

	database.DB.First(&abrigo, id)

	if abrigo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Abrigo nao encontrado",
		})
		return
	}

	database.DB.Delete(&abrigo, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Abrigo deletado com sucesso!",
	})
}
