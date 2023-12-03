package middleware

import (
	"api/database"
	"api/models"
	"fmt"
	"os"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// функция записи данных авторизированного клиента в cookies
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("middleware")

		tokenString, err := c.Cookie("Authorization")
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("error: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("SECRET")), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			var client1 models.Client
			database.GormDB.First(&client1, claims["sub"])

			if client1.Client_Id == 0 {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			c.Set("client", client1.Client_Login)

			c.Next()
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
