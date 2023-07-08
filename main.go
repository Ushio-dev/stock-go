package main

import (
	"github.com/Ushio-dev/stock-go/controllers"
	"github.com/Ushio-dev/stock-go/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/nuevo", controllers.ItemCreate)
	r.GET("/items", controllers.ReadItems)
	r.GET("/item/:id", controllers.ReadOneItem)
	r.PUT("/item/actualizar/:id", controllers.Updateitem)
	r.Run()
}
