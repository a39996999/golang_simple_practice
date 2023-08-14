package user

import (
	"chatroom/model"
	"chatroom/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name     string `json:"Username"`
	Passowrd string `json:"Password"`
	Email    string `json:"Email"`
}

func CreateUser(c *gin.Context) {
	User := User{}
	if err := c.ShouldBindJSON(&User); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  "10001",
			"error": err.Error(),
		})
		return
	}
	if User.Name == "" || User.Passowrd == "" || User.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  "10001",
			"error": "invalid input",
		})
		return
	}
	checkEmailValidate, emailError := utils.VerifyEmailFormat(User.Email)
	if emailError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  "10001",
			"error": emailError,
		})
	} else if checkEmailValidate == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  "10001",
			"error": "Input wrong email format",
		})
	} else {
		err := model.CreateUser(User.Name, User.Passowrd, User.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":  "10001",
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":   "0",
				"status": "create user sucessfully",
			})
		}
	}
}

func UpdateUserPassword(c *gin.Context) {
	User := User{}
	if err := c.ShouldBindJSON(&User); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  "10001",
			"error": err.Error(),
		})
	}
	if User.Name == "" || User.Passowrd == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  "10001",
			"error": "invalid input",
		})
	}
	err := model.UpdateUserPassword(User.Name, User.Passowrd)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  "10001",
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":   "0",
			"status": "update user password sucessfully",
		})
	}
}

func DeleteUser(c *gin.Context) {
	User := User{}
	if err := c.ShouldBindJSON(&User); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  "10001",
			"error": err.Error(),
		})
	}
	if User.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  "10001",
			"error": "invalid input",
		})
	}
	err := model.DeleteUser(User.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  "10001",
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   "0",
		"status": "delete user sucessfully",
	})
}

func UserLogin(c *gin.Context) {
	User := User{}
	if err := c.ShouldBindJSON(&User); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  "10001",
			"error": err.Error(),
		})
		return
	}
	passwordHash, salt, err := model.QueryPassword(User.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  "10001",
			"error": err.Error(),
		})
		return
	}
	userPasswordHash := utils.HashPassword(User.Passowrd, salt)
	if passwordHash == userPasswordHash {
		c.JSON(http.StatusOK, gin.H{
			"code":    "0",
			"message": "Login sucessfully",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":  "10001",
			"error": "input wrong password",
		})
	}
}
