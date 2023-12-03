package models

//структура запроса на хранение оборудования
type EquipmentKeepingRequests struct {
	Request_Id     uint   `json:"request_id" validate:"required"`
	Equipment_Code string `json:"equipment_code" validate:"required"`
	Name_Status    string `json:"name_status" validate:"required"`
}
