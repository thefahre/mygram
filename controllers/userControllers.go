package controllers

import (
	"mygram/database"
	"mygram/helpers"
	"mygram/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

// UserRegister godoc
// @Summary Register a new user
// @Description Register a new user via endpoint
// @Tags user
// @Accept json
// @Produce json
// @Param payload body requests.UserRegisterRequest true "New User Data"
// @Success 201 {object} requests.UserRegisterResponse
// @Router /users/register [post]
func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}
	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":       User.ID,
		"username": User.Username,
		"email":    User.Email,
		"age":      User.Age,
	})
}

// UserLogingodoc
// @Summary User login
// @Description User login to obtain JWT Beraer token
// @Tags user
// @Accept json
// @Produce json
// @Param payload body requests.UserLoginRequest true "Login Data"
// @Success 200 {object} requests.UserLoginResponse
// @Router /users/login [post]
func UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}
	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Not Authorized",
			"message": "invalid credentials",
		})
		return
	}
	// //debug!
	// fmt.Println("input Pasword: ", helpers.HashPass(password))
	// fmt.Println("User Pasword: ", User.Password)
	// //debug!
	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Not Authorized",
			"message": "invalid credentials",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
