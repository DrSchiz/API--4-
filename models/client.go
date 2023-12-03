package models

//структура клиента
type Client struct {
	Client_Id           uint   `gorm:"primaryKey"`
	Client_Login        string `json:"login" validate:"required,min=3,max=20"`
	Client_Hash         string `json:"hash" validate:"required"`
	Client_Name         string `json:"name" validate:"required"`
	Client_Firstname    string `json:"firstname" validate:"required"`
	Client_Patronymic   string `json:"patronymic"`
	Client_Phone_Number string `json:"phone" validate:"required,numeric"`
	Client_Email        string `json:"email" validate:"required,email"`
}
