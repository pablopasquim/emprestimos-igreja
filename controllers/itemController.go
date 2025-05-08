package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pablopasquim/emprestimos-igreja/models"
)

type itemController struct {
}

func NewItemController() itemController {
	return itemController{}
}

func (i *itemController) GetItem(ctx *gin.Context) {

	item := []models.Item{
		{
			ID:       1,
			Name:     "Cadeira",
			Category: "Assento",
			Status:   "Disponível",
		},
	}

	ctx.JSON(http.StatusOK, item)
}
