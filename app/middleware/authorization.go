package middleware

import (
	"jwt-go/app/entity"
	"jwt-go/pkg/database"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ProductAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, _ := database.Connect()
		productID, err := strconv.Atoi(c.Param("productId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Unauthorized",
				"message": "Invalid parameters",
			})
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		userLevel := userData["level"]
		Product := entity.Product{}

		err = db.Select("user_id").First(&Product, uint(productID)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "Data doesn't exist",
			})
			return
		}

		//kalau level user dan Product.UserID != userID maka error
		if Product.UserID != userID && userLevel == "user" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
		}
	}
}
