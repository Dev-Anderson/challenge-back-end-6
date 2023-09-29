package controllers

import (
	"adopet/database"
	"adopet/models"
	"adopet/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Criar um user
// @Summary Criar um novo usuario
// @Description Cria um novo usuario
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.User true "Dados do usuario"
// @Success 201
// @Failure 404 {object} error
// @Router /user [post]
func CreateUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Nao foi possivel vincular o JSON" + err.Error(),
		})
		return
	}

	user.Password = services.Sha256Encoder(user.Password)

	err = database.DB.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Nao foi possivel criar o usuario" + err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}

// Consulta de todos os user
// @Summary Consulta todos os usuarios
// @Description Rota para consultar todos os user (rota nao deve ir para producao)
// @Schemes
// @Tags User
// @Accet json
// @Produce json
// @Success 200 {object} models.User
// @Router /user [get]
func GetAllUser(c *gin.Context) {
	var user []models.User
	database.DB.Find(&user)
	c.JSON(http.StatusOK, user)
}
