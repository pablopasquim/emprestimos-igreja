package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pablopasquim/emprestimos-igreja/database"
	"github.com/pablopasquim/emprestimos-igreja/models"
)

func CreateUserAdmin(ctx *gin.Context) {
	var userAdmin models.UserAdmin

	database.DB.Create(&userAdmin)
}
