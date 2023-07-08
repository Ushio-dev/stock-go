package controllers

import (
	"github.com/Ushio-dev/stock-go/initializers"
	"github.com/Ushio-dev/stock-go/models"
	"github.com/gin-gonic/gin"
)

func ItemCreate(c *gin.Context) {
	var newItem struct {
		Name   string
		Amount uint
	}
	c.Bind(&newItem)

	item := models.Item{Name: newItem.Name, Amount: newItem.Amount}
	result := initializers.DB.Create(&item)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"item": item,
	})
}

func ReadItems(c *gin.Context) {
	var items []models.Item
	initializers.DB.Find(&items)

	c.JSON(200, gin.H{
		"items": items,
	})
}

func ReadOneItem(c *gin.Context) {
	id := c.Param("id")

	var item models.Item

	result := initializers.DB.First(&item, id)

	if result.Error != nil {
		c.Status(404)
		return
	}
	c.JSON(200, gin.H{
		"item": item,
	})
}

func Updateitem(c *gin.Context) {
	id := c.Param("id")

	var itemDto struct {
		Name   string
		Amount uint
	}

	c.Bind(&itemDto)
	var item models.Item
	result := initializers.DB.First(&item, id)

	if result.Error != nil {
		c.Status(404)
		return
	}

	initializers.DB.Model(&item).Updates(models.Item{
		Name:   itemDto.Name,
		Amount: itemDto.Amount,
	})

	c.JSON(200, gin.H{
		"item": item,
	})
}
