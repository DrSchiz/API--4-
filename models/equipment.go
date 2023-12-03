package models

//структура оборудования
type Equipments struct {
	Equipment_Code   string `json:"equipment_code" validate:"required"`
	Equipment_Size   int    `json:"equipment_size" validate:"required"`
	Type_Name        string `json:"type_name" validate:"required"`
	Warehouse_Number *int   `json:"warehouse_number"`
	Status_Name      string `json:"status_name" validate:"required"`
	Client_Login     string `json:"client_login" validate:"required"`
	Status_Keep      *bool  `json:"status_keep" validate:"required"`
	Equipment_Id     *uint
	Status_Request   *string `json:"status_request"`
}
