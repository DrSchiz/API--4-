package models

//струкрута вида оборудования
type Type_Equipment struct {
	Type_Name     string `json:"type_name" validate:"required"`
	Category_Name string `json:"category_name" validate:"required"`
	Type_Id       uint   `gorm:"primaryKey"`
}
