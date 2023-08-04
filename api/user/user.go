package user

import (
	"chatroom/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name     string `json:"Username"`
	Passowrd string `json:"Password"`
}

func CreateUser(c *gin.Context) {
	User := User{}
	if err := c.ShouldBindJSON(&User); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
	}
	checkUserExist := model.FindUserExist(User.Name)
	if checkUserExist {
		c.JSON(http.StatusOK, gin.H{"error": "User is alreay exist"})
	} else {
		err := model.CreateUser(User.Name, User.Passowrd)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "insert db error"})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": "create user sucessfully"})
		}
	}
}

func UpdateUserProfile(c *gin.Context) {
	User := User{}
	if err := c.ShouldBindJSON(&User); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
	}
	checkUserExist := model.FindUserExist(User.Name)
	if checkUserExist {
		err := model.UpdateUserPassword(User.Name, User.Passowrd)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "sql update: error "})
		}
		c.JSON(http.StatusOK, gin.H{"status": "update user password sucessfully"})
	} else {
		c.JSON(http.StatusOK, gin.H{"error": "User is not exist"})
	}
}

func DeleteUser(c *gin.Context) {
	User := User{}
	if err := c.ShouldBindJSON(&User); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
	}
	checkUserExist := model.FindUserExist(User.Name)
	if checkUserExist {
		err := model.DeleteUser(User.Name)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "sql update: error "})
		}
		c.JSON(http.StatusOK, gin.H{"status": "delete user sucessfully"})
	} else {
		c.JSON(http.StatusOK, gin.H{"error": "User is not exist"})
	}
}
