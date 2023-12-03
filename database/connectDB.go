package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// инициализация переменной БД
var GormDB *gorm.DB

// функция подключения к БД
func Connect() {
	var err error
	GormDB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=postgres password=postgres dbname=IEAIS port=5432 sslmode=disable",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("База данных успешно подключена!")
}
