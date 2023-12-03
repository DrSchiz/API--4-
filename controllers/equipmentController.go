package controllers

import (
	"api/database"
	"api/http/equipment"
	"api/models"
	"encoding/json"

	"github.com/gin-gonic/gin"
	// "github.com/go-playground/validator/v10"
)

// функция получения клиентского оборудования по его номеру
func GetEquipment(c *gin.Context) {
	_client := Validate(c)

	var equipment models.Equipments

	database.GormDB.Where("equipment_code = ?", c.Param("code_equipment")).Find(&equipment)

	if equipment.Client_Login != _client.Client_Login {
		c.JSON(400, gin.H{
			"error": "Данное оборудование вам не принадлежит!",
		})
		return
	}

	c.JSON(200, equipment)
}

// функция получения оборудования авторизированного клиента
func GetEquipments(c *gin.Context) {
	_client := Validate(c)

	var equipments []models.Equipments

	database.GormDB.Where("client_login = ?", _client.Client_Login).Find(&equipments)

	c.JSON(200, equipments)
}

// функция получения статусов целостности оборудования
func GetEquipmentStatus(c *gin.Context) {
	var equipmentStatuses []models.Equipment_Status

	database.GormDB.Find(&equipmentStatuses)

	c.JSON(200, equipmentStatuses)
}

// функция получения видов оборудования
func GetTypeEquipment(c *gin.Context) {
	var typeEquipments []models.Type_Equipment

	database.GormDB.Find(&typeEquipments)

	c.JSON(200, typeEquipments)
}

// функция получения заявок на хранение оборудования
func GetEquipmentKeepingRequests(c *gin.Context) {
	var equipmentKeepingRequests []models.EquipmentKeepingRequests

	database.GormDB.Find(&equipmentKeepingRequests)

	c.JSON(200, equipmentKeepingRequests)
}

// функция создания заявки на хранение оборудования
func CreateKeepingRequest(c *gin.Context) {

	_client := Validate(c)

	var equipments []models.Equipments

	database.GormDB.Where("client_login = ?", _client.Client_Login).Find(&equipments)

	for _, i := range equipments {
		if i.Equipment_Code == c.Param("code_equipment") {

			err := database.GormDB.Exec("CALL public.create_request(?)", c.Param("code_equipment"))
			if err.Error != nil {
				c.JSON(400, gin.H{
					"error": "Ошибка при создании заявки!",
				})
				return
			}

			c.JSON(200, gin.H{
				"success": "Заявка успешно подана",
			})
			return
		}
	}

	c.JSON(400, gin.H{
		"error": "Данное оборудование вам не принадлежит!",
	})
}

// функция создания оборудования
func CreateEquipment(c *gin.Context) {

	_client := Validate(c)

	var equipments []models.Equipments

	database.GormDB.Where("client_login = ?", _client.Client_Login).Find(&equipments)

	var newEquipment equipment.RequestCreateEquipment

	json.NewDecoder(c.Request.Body).Decode(&newEquipment)

	var _equipment models.Equipments

	_equipment = models.Equipments{
		Equipment_Code:   newEquipment.Equipment_Code,
		Equipment_Size:   newEquipment.Equipment_Size,
		Type_Name:        newEquipment.Type_Name,
		Status_Name:      newEquipment.Status_Name,
		Client_Login:     _client.Client_Login,
		Warehouse_Number: nil,
	}

	create := database.GormDB.Create(_equipment)
	if create.Error != nil {
		c.JSON(400, gin.H{
			"error": create.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"success": "Оборудование успешно добавлено!",
	})
}

// функция удаления оборудования
func DeleteEquipment(c *gin.Context) {
	_client := Validate(c)

	var equipment models.Equipments

	database.GormDB.Where("equipment_code = ?", c.Param("code_equipment")).Find(&equipment)

	if equipment.Client_Login != _client.Client_Login {
		c.JSON(400, gin.H{
			"error": "Оборудование вам не принадлежит!",
		})
		return
	}

	database.GormDB.Exec("delete from equipments where equipment_code = ?", equipment.Equipment_Code)

	c.JSON(200, gin.H{
		"success": "Оборудование успешно удалено!",
	})
}

// функция изменения данных оборудования
func EditEquipment(c *gin.Context) {
	_client := Validate(c)

	var _equipment models.Equipments

	database.GormDB.Where("equipment_code = ?", c.Param("code_equipment")).Find(&_equipment)

	if _equipment.Client_Login != _client.Client_Login {
		c.JSON(400, gin.H{
			"error": "Оборудование вам не принадлежит!",
		})
		return
	}

	var editEquipment equipment.RequestCreateEquipment

	json.NewDecoder(c.Request.Body).Decode(&editEquipment)

	_equipment.Equipment_Code = editEquipment.Equipment_Code
	_equipment.Equipment_Size = editEquipment.Equipment_Size
	_equipment.Type_Name = editEquipment.Type_Name
	_equipment.Status_Name = editEquipment.Status_Name

	database.GormDB.Where("equipment_code = ?", c.Param("code_equipment")).Save(&_equipment)

	c.JSON(200, gin.H{
		"success": "Данные оборудования успешно обновлены",
	})
}
