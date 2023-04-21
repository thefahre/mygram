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

// CreatePhoto godoc
// @Summary Create one new photo entry
// @Description Create one new photo entry with logged in user
// @Tags photo
// @Accept json
// @Produce json
// @Param payload body requests.CreatePhotoRequest true "New Photo Data"
// @Success 200 {object} models.Photo
// @Router /photos/ [post]
// @Security Bearer
func CreatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID

	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, Photo)
}

// UpdatePhoto godoc
// @Summary Update photo entry
// @Description Update photo entry with logged in user (authorized user only)
// @Tags photo
// @Accept json
// @Produce json
// @Param id path int true "ID of photo to be updated"
// @Param payload body requests.CreatePhotoRequest true "New Photo Data"
// @Success 200 {object} models.Photo
// @Router /photos/{id} [put]
// @Security Bearer
func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	Photo.ID = uint(photoId)

	err := db.Model(&Photo).Where("id = ?", photoId).Updates(models.Photo{
		Title:   Photo.Title,
		Caption: Photo.Caption,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Photo)
}

// GetAllPhoto godoc
// @Summary Get all photo
// @Description Get all photos with currently logged in user
// @Tags photo
// @Accept json
// @Produce json
// @Success 200 {object} models.Photo
// @Router /photos/ [get]
// @Security Bearer
func GetAllPhoto(c *gin.Context) {
	db := database.GetDB()
	photos := []models.Photo{}

	if err := db.Preload("Comments").Find(&photos).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "unable to find any photo",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": photos,
	})
}

// GetOnePhoto godoc
// @Summary Get one photo
// @Description Get one photo with currently logged in user
// @Tags photo
// @Accept json
// @Produce json
// @Param id path int true "Photo ID"
// @Success 200 {object} models.Photo
// @Router /photos/{id} [get]
// @Security Bearer
func GetOnePhoto(c *gin.Context) {
	db := database.GetDB()
	// userData := c.MustGet("userData").(jwt.MapClaims)
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	photo := []models.Photo{}

	if err := db.Preload("Comments").Where("id = ?", photoId).First(&photo).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "photo id not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": photo,
	})
}

// DeletePhoto godoc
// @Summary Delete photo entry
// @Description Delete photo entry with logged in user (authorized user only)
// @Tags photo
// @Accept json
// @Produce json
// @Param id path int true "ID of photo to be deleted"
// @Success 204 "No Content"
// @Router /photos/{id} [delete]
// @Security Bearer
func DeletePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&photo)
	} else {
		c.ShouldBind(&photo)
	}

	photo.UserID = userID
	photo.ID = uint(photoId)

	if err := db.Where("id = ?", photoId).Delete(&photo).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "photo id not found.",
		})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
