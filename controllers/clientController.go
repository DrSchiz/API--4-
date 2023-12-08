package controllers

import (
	"api/database"
	"api/http/client"
	"api/models"
	"encoding/json"
	"os"
	"time"

	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/lib/pq"
)

// инициализация переменной валиданции данных
var Validator = validator.New()

// функция хэширования пароля
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// функция проверки хэша
func CheckHashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// функция получения клиента по его логину
func GetClientByLogin(c *gin.Context, login string) models.Client {
	var foundClient models.Client

	database.GormDB.Where("client_login = ?", login).Find(&foundClient)

	return foundClient
}

// функция получения авторизированного клиента по логину
func ShowClient(c *gin.Context) {
	foundClient := GetClientByLogin(c, c.Param("login"))
	c.JSON(200, foundClient)
}

// функция регистрации
func CreateClient(c *gin.Context) {
	var createClient client.RequestCreateClient

	json.NewDecoder(c.Request.Body).Decode(&createClient)

	hash, _ := HashPassword(createClient.Client_Password)

	clientCreate := &models.Client{
		Client_Login:        createClient.Client_Login,
		Client_Hash:         hash,
		Client_Name:         createClient.Client_Name,
		Client_Firstname:    createClient.Client_Firstname,
		Client_Patronymic:   createClient.Client_Patronymic,
		Client_Phone_Number: createClient.Client_Phone_Number,
		Client_Email:        createClient.Client_Email,
	}

	err := Validator.Struct(clientCreate)
	if err != nil {
		if createClient.Client_Password != createClient.Client_Repeat_Password {
			c.JSON(403, gin.H{
				"error": "Пароли не совпадают!",
			})
		}

		if createClient.Client_Password == "" {
			c.JSON(403, gin.H{
				"error": "Пароль не может быть пустым!",
			})
		}

		if len(createClient.Client_Password) <= 3 || len(createClient.Client_Password) >= 50 {
			c.JSON(403, gin.H{
				"error": "Пароль должен содержать от 3 до 50 символов!",
			})
		}
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	create := database.GormDB.Create(clientCreate)
	if create.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": create.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"success": "Аккаунт успешно создан!",
	})
}

// функция изменения клиентских данных
func UpdateClientData(c *gin.Context) {
	clientUpdate := Validate(c)

	var updateClient client.RequestUpdateClient
	json.NewDecoder(c.Request.Body).Decode(&updateClient)

	update := database.GormDB.Model(&clientUpdate).Where("client_login = ?", clientUpdate.Client_Login).Updates(updateClient)
	if update.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": update.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"success": "Данные аккаунта успешно обновлены!",
	})
}

// функция авторизации клиента
func LoginClient(c *gin.Context) {
	var clientLogin client.RequestLoginClient
	json.NewDecoder(c.Request.Body).Decode(&clientLogin)

	loginClient := GetClientByLogin(c, clientLogin.Client_Login)

	if !CheckHashPassword(clientLogin.Client_Password, loginClient.Client_Hash) {
		c.JSON(400, gin.H{
			"error": "Неверный логин или пароль!",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": loginClient.Client_Id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(403, gin.H{
			"error": "Ошибка создания токена!",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(200, gin.H{
		"token": tokenString,
	})
}

// функция валидации данных
func Validate(c *gin.Context) models.Client {

	_client, _ := c.Get("client")

	var client1 models.Client
	client1 = GetClientByLogin(c, _client.(string))

	return client1
}

// функция получения авторизированного клиента
func GetAuthorizedClient(c *gin.Context) {
	_client := Validate(c)

	if _client == (models.Client{}) {
		c.JSON(401, gin.H{"error": "Пользователь не авторизирован!"})
		return
	}

	c.JSON(200, _client)
}

// функция смены пароля авторизированного клиента
func ChangePassword(c *gin.Context) {

	var clientData client.RequestChangePassword
	json.NewDecoder(c.Request.Body).Decode(&clientData)

	_client := Validate(c)

	if !CheckHashPassword(clientData.Client_Old_Password, _client.Client_Hash) {
		c.JSON(400, gin.H{
			"error": "Неверный пароль!",
		})
		return
	}

	if clientData.Client_New_Password != clientData.Client_Repeat_Password {
		c.JSON(400, gin.H{
			"error": "Пароли не совпадают!",
		})
		return
	}

	hash, _ := HashPassword(clientData.Client_New_Password)
	_client.Client_Hash = hash

	database.GormDB.Save(&_client).Where("client_login = ?", _client.Client_Login)

	c.JSON(200, gin.H{
		"success": "Пароль успешно обновлён!",
	})
}
