package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/fbpr/task-5-vix-btpns-febry-prasetya/database"
	"github.com/fbpr/task-5-vix-btpns-febry-prasetya/helpers"
	"github.com/fbpr/task-5-vix-btpns-febry-prasetya/models"
	"github.com/gin-gonic/gin"
)

func CreatePhoto(c *gin.Context) {
	db := database.GetDb()

	photo := models.Photo{}
	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userEmail := c.MustGet("userData").(*helpers.JwtCustomClaims).Email
	user := models.User{}
	err := db.Debug().Where("email = ?", userEmail).First(&user).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	photo.UserID = user.ID
	err = db.Debug().Create(&photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "photo successfully created", "data": gin.H{"title": photo.Title,
		"caption": photo.Caption, "photo_url": photo.PhotoUrl}})
}

func GetPhoto(c *gin.Context) {
	db := database.GetDb()
	
	photos := []map[string]interface{}{}

	if err := db.Debug().Model(models.Photo{}).Find(&photos).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": photos})
}

func UpdatePhoto(c *gin.Context) {
	db := database.GetDb()

	photo := models.Photo{}
	newPhoto := models.Photo{}

	err := db.First(&photo, c.Param("userId")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err,})
		return
	}

	err = json.NewDecoder(c.Request.Body).Decode(&newPhoto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "update success"})
}

func DeletePhoto(c *gin.Context) {
	db := database.GetDb()
	photo := models.Photo{}

	err := db.First(&photo, c.Param("photoId")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = db.Debug().Delete(&photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "photo deleted"})
}
