package controllers

import (
	"mygram/database"
	"mygram/helpers"
	"mygram/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func CreateSocialmedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Socialmedia := models.Socialmedia{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Socialmedia)
	} else {
		c.ShouldBind(&Socialmedia)
	}

	Socialmedia.UserID = userID

	err := db.Debug().Create(&Socialmedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, Socialmedia)
}

func UpdateSocialmedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Socialmedia := models.Socialmedia{}

	socialmediaId, _ := strconv.Atoi(c.Param("socialmediaId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Socialmedia)
	} else {
		c.ShouldBind(&Socialmedia)
	}

	Socialmedia.UserID = userID
	Socialmedia.ID = uint(socialmediaId)

	err := db.Model(&Socialmedia).Where("id = ?", socialmediaId).Updates(models.Socialmedia{
		Name:           Socialmedia.Name,
		SocialMediaURL: Socialmedia.SocialMediaURL,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Socialmedia)
}

func GetAllSocialmedia(c *gin.Context) {
	db := database.GetDB()
	socialmedias := []models.Socialmedia{}

	if err := db.Find(&socialmedias).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "unable to find any social media data",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": socialmedias,
	})
}

func GetOneSocialmedia(c *gin.Context) {
	db := database.GetDB()
	// userData := c.MustGet("userData").(jwt.MapClaims)
	socialmediaId, _ := strconv.Atoi(c.Param("socialmediaId"))
	socialmedia := []models.Socialmedia{}

	if err := db.Where("id = ?", socialmediaId).First(&socialmedia).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "social media id not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": socialmedia,
	})
}

func DeleteSocialmedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	socialmedia := models.Socialmedia{}

	socialmediaId, _ := strconv.Atoi(c.Param("socialmediaId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&socialmedia)
	} else {
		c.ShouldBind(&socialmedia)
	}

	socialmedia.UserID = userID
	socialmedia.ID = uint(socialmediaId)

	if err := db.Where("id = ?", socialmediaId).Delete(&socialmedia).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "social media id not found.",
		})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
