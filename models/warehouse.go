package models

//структура склада
type Warehouse struct {
	Warehouse_Number   uint `gorm:"primaryKey"`
	Warehouse_Capacity int  `json:"warehouse_capacity" validate:"required"`
	Warehouse_Fullness int  `json:"warehouse_fullness" validate:"required"`
}
