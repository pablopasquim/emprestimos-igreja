package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pablopasquim/emprestimos-igreja/controllers"
	"github.com/pablopasquim/emprestimos-igreja/database"
)

func main() {
	server := gin.Default()

	database.ConectionDataBase()

	ItemController := controllers.NewItemController()

	server.GET("/itens", ItemController.GetItem)

	server.Run(":8080")
}
