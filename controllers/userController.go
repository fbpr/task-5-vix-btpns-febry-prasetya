package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/fbpr/task-5-vix-btpns-febry-prasetya/database"
	"github.com/fbpr/task-5-vix-btpns-febry-prasetya/helpers"
	"github.com/fbpr/task-5-vix-btpns-febry-prasetya/models"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	db := database.GetDb()
	user := models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := db.Debug().Model(models.Photo{}).Create(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "registration success", "data": gin.H{"username": user.Username, "email": user.Email}})
}

func UpdateUserById(c *gin.Context) {
	db := database.GetDb()

	user := models.User{}
	newUser := models.User{}

	err := db.First(&user, c.Param("userId")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = json.NewDecoder(c.Request.Body).Decode(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = db.Debug().Model(&user).Updates(newUser).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "update user success"})
}

func DeleteUserById(c *gin.Context) {
	db := database.GetDb()

	user := models.User{}
	err := db.First(&user, c.Param("userId")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = db.Debug().Delete(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}

func Login(c *gin.Context) {
	db := database.GetDb()
	user := models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	password := user.Password

	err := db.Debug().Where("email = ?", user.Email).Take(&user).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	comparePassword := helpers.VerifyPassword(user.Password, password)
	if !comparePassword {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := helpers.GenerateJWT(user.Username, user.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "login success", "token": token})
}
