package migrate

import (
	"api/database"
	"api/models"
)

// функция миграции таблиц из БД
func Migrate() {
	database.GormDB.AutoMigrate(&models.Client{})
	database.GormDB.AutoMigrate(&models.Equipments{})
	database.GormDB.AutoMigrate(&models.Equipment_Status{})
	database.GormDB.AutoMigrate(&models.Type_Equipment{})
	database.GormDB.AutoMigrate(&models.Warehouse{})
}
