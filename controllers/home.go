package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ShowHome godoc
// @Summary      Rota para testar a API
// @Description  Rota para testar a API
// @Tags         home
// @Produce      json
// @Success      200  {object}  string
// @Router       /home	[get]
func GetHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "API rodando com sucesso",
	})
}
