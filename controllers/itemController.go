package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pablopasquim/emprestimos-igreja/models"
	"net/http"
)

type itemController struct {
}

func NewItemController() itemController {
	return itemController{}
}

func (i *itemController) GetItem(ctx *gin.Context) {

	item := []models.Item{
		{
			ID:     1,
			Nome:   "Cadeira",
			Tipo:   "Assento",
			Status: "Disponível",
		},
	}

	ctx.JSON(http.StatusOK, item)
}
