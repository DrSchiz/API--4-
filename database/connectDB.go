package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// инициализация переменной БД
var GormDB *gorm.DB

// функция подключения к БД
func Connect() {
	var err error
	GormDB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=" + os.Getenv("USER") + " password=" + os.Getenv("PASSWORD") + " dbname=" + os.Getenv("DBNAME") + " port=" + os.Getenv("PORT") + " sslmode=disable",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("База данных успешно подключена!")
}
