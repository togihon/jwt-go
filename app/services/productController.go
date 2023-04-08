package services

import (
	"jwt-go/app/entity"
	"jwt-go/pkg/database"
	"jwt-go/pkg/helper"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ReadProduct(c *gin.Context) {
	db, _ := database.Connect()
	contentType := helper.GetContentType(c)
	Product := entity.Product{}

	//get parameter
	productID, _ := strconv.Atoi(c.Param("productId"))

	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	//query select * from product where id = param
	err := db.First(&Product, "id = ?", productID).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Product)
}

func CreateProduct(c *gin.Context) {
	db, _ := database.Connect()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helper.GetContentType(c)

	Product := entity.Product{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	Product.UserID = userID
	err := db.Debug().Create(&Product).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Product)
}

func UpdateProduct(c *gin.Context) {
	db, _ := database.Connect()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helper.GetContentType(c)
	Product := entity.Product{}
	userLevel := userData["level"]

	if userLevel == "user" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "You are not allowed to access this data",
		})
		return
	}

	productID, _ := strconv.Atoi(c.Param("productId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	Product.UserID = userID
	Product.ID = uint(productID)

	err := db.Model(&Product).Where("id = ?", productID).Updates(entity.Product{Title: Product.Title, Description: Product.Description}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Product)
}

func DeleteProduct(c *gin.Context) {
	db, _ := database.Connect()
	contentType := helper.GetContentType(c)
	Product := entity.Product{}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userLevel := userData["level"]

	if userLevel == "user" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "You are not allowed to access this data",
		})
		return
	}

	//get parameter
	productID, _ := strconv.Atoi(c.Param("productId"))

	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	err := db.Where("id = ?", productID).Delete(&Product).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data berhasil dihapus",
	})
}
