package controllers

import (
	"api/database"
	"api/models"

	"github.com/gin-gonic/gin"
	// "github.com/go-playground/validator/v10"
)

// функция получения складов
func GetWarehouses(c *gin.Context) {
	var warehouses []models.Warehouse

	database.GormDB.Find(&warehouses)

	c.JSON(200, warehouses)
}
