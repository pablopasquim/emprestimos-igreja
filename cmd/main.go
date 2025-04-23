package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pablopasquim/emprestimos-igreja/controllers"
)

func main() {
	server := gin.Default()

	ItemController := controllers.NewItemController()

	server.GET("/itens", ItemController.GetItem)

	server.Run(":8080")
}
