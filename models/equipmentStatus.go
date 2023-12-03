package models

//структура статуса целостности оборудования
type Equipment_Status struct {
	Status_Id   uint   `gorm:"primaryKey"`
	Status_Name string `json:"status_name" validate:"required"`
}
