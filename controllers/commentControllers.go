package controllers

import (
	"fmt"
	"mygram/database"
	"mygram/helpers"
	"mygram/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// CreateComment godoc
// @Summary Create new comment
// @Description Create new comment on one photo by id
// @Tags comment
// @Accept json
// @Produce json
// @Param photo_id path int true "ID of the photo to add comment to"
// @Param payload body requests.CreateCommentRequest true "New comment data"
// @Success 200 {object} models.Comment
// @Router /photos/{photo_id}/comments [post]
// @Security Bearer
func CreateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Comment := models.Comment{}
	userID := uint(userData["id"].(float64))
	photoID, _ := strconv.Atoi(c.Param("photoId"))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	Comment.PhotoID = uint(photoID)

	err := db.Debug().Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, Comment)
}

// UpdateComment godoc
// @Summary Update a comment
// @Description Update a comment on one photo by both ids
// @Tags comment
// @Accept json
// @Produce json
// @Param photo_id path int true "ID of the photo to update comment to"
// @Param comment_id path int true "ID of the comment to update"
// @Param payload body requests.CreateCommentRequest true "New comment data"
// @Success 200 {object} models.Comment
// @Router /photos/{photo_id}/comments/{comment_id} [put]
// @Security Bearer
func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	Comment.PhotoID = uint(photoId)
	Comment.ID = uint(commentId)

	err := db.Model(&Comment).Where("id = ?", commentId).Where("photo_id = ?", photoId).Updates(models.Comment{
		Message: Comment.Message,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Comment)
}

// GetAllComment godoc
// @Summary Get all comments from one photo
// @Description Get all comments from one photo with currently logged in user
// @Tags comment
// @Accept json
// @Produce json
// @Param photo_id path int true "ID of the photo to fetch comments from"
// @Success 200 {object} models.Comment
// @Router /photos/{photo_id}/comments [get]
// @Security Bearer
func GetAllComment(c *gin.Context) {
	db := database.GetDB()
	comments := []models.Comment{}
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	fmt.Println(photoId)
	if err := db.Where("photo_id = ?", photoId).Find(&comments).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "unable to find any comments with this photo id",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": comments,
	})
}

// GetOneComment godoc
// @Summary Get one comment from one photo
// @Description Get one comment by comment ID from one photo with currently logged in user
// @Tags comment
// @Accept json
// @Produce json
// @Param photo_id path int true "ID of the photo to fetch comments from"
// @Param comment_id path int true "ID of the comment to fetch"
// @Success 200 {object} models.Comment
// @Router /photos/{photo_id}/comments/{comment_id} [get]
// @Security Bearer
func GetOneComment(c *gin.Context) {
	db := database.GetDB()
	// userData := c.MustGet("userData").(jwt.MapClaims)
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	Comment := []models.Comment{}

	if err := db.Where("photo_id = ?", photoId).Where("id = ?", commentId).First(&Comment).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "comment or photo id not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": Comment,
	})
}

// DeleteComment godoc
// @Summary Delete comment
// @Description Delete comment from a photo with logged in user (authorized user only)
// @Tags comment
// @Accept json
// @Produce json
// @Param photo_id path int true "ID of photo to delete comment from"
// @Param comment_id path int true "ID of comment to delete"
// @Success 204 "No Content"
// @Router /photos/{photo_id}/comments/{comment_id} [delete]
// @Security Bearer
func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	Comment.ID = uint(commentId)
	Comment.PhotoID = uint(photoId)

	if err := db.Where("photo_id = ?", photoId).Where("id = ?", commentId).Delete(&Comment).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "comment id not found.",
		})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
