package main

import (
	"github.com/Ushio-dev/stock-go/initializers"
	"github.com/Ushio-dev/stock-go/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Item{})
}
