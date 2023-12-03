package main

import (
	"api/controllers"
	"api/database"
	"api/middleware"
	"api/migrate"

	"github.com/gin-gonic/gin"
)

// инициализация подключения к БД и миграции таблиц из БД
func init() {
	database.Connect()
	migrate.Migrate()
}

// запуск API и указание маршрутов для работы с ним
func main() {
	r := gin.Default()

	authorized := r.Group("/main")

	authorized.Use(middleware.Logger())

	r.POST("/create-account", controllers.CreateClient)
	r.GET("/:login", controllers.ShowClient)
	r.POST("/login", controllers.LoginClient)
	r.GET("/equipment_statuses", controllers.GetEquipmentStatus)
	r.GET("/type_equipments", controllers.GetTypeEquipment)
	r.GET("/warehouses", controllers.GetWarehouses)
	r.GET("/equipment-keeping-requests", controllers.GetEquipmentKeepingRequests)

	authorized.GET("/validate", controllers.GetAuthorizedClient)
	authorized.PUT("/client", controllers.UpdateClientData)
	authorized.PUT("/client/password", controllers.ChangePassword)
	authorized.GET("/equipment/:code_equipment", controllers.GetEquipment)
	authorized.GET("/equipment", controllers.GetEquipments)
	authorized.POST("/equipment/:code_equipment/create-request", controllers.CreateKeepingRequest)
	authorized.POST("/equipment/create-equipment", controllers.CreateEquipment)
	authorized.DELETE("/equipment/:code_equipment/delete", controllers.DeleteEquipment)
	authorized.PUT("/equipment/:code_equipment/edit", controllers.EditEquipment)

	r.Run()
}
