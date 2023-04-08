package services

import (
	"jwt-go/app/entity"
	"jwt-go/pkg/database"
	"jwt-go/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func UserRegister(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	contentType := helper.GetContentType(c)
	_, _ = db, contentType
	User := entity.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err = db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":        User.ID,
		"email":     User.Email,
		"full_name": User.FullName,
		"level":     User.Level,
	})

}

func UserLogin(c *gin.Context) {
	db, _ := database.Connect()
	contentType := helper.GetContentType(c)
	_, _ = db, contentType

	User := entity.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	//select data user berdasarkan email
	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email or password",
		})
		return
	}

	comparePass := helper.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email or password",
		})
		return
	}

	token := helper.GenerateToken(User.ID, User.Email, User.Level)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
