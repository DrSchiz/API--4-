package equipment

//структура запроса создания оборудования
type RequestCreateEquipment struct {
	Equipment_Code string `json:"equipment_code" validate:"required"`
	Type_Name      string `json:"type_name" validate:"required"`
	Equipment_Size int    `json:"equipment_size" validate:"required"`
	Status_Name    string `json:"status_name" validate:"required"`
}
