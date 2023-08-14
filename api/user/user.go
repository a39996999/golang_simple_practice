package user

import (
	"chatroom/middleware"
	"chatroom/model"
	"chatroom/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name     string `json:"Username"`
	Password string `json:"Password"`
	Email    string `json:"Email"`
}

func CreateUser(c *gin.Context) {
	User := User{}
	if err := c.ShouldBindJSON(&User); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if User.Name == "" || User.Password == "" || User.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input",
		})
		return
	}
	checkEmailValidate, emailError := utils.VerifyEmailFormat(User.Email)
	if emailError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": emailError,
		})
	} else if checkEmailValidate == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Input wrong email format",
		})
	} else {
		err := model.CreateUser(User.Name, User.Password, User.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "create user sucessfully",
			})
		}
	}
}

func UpdateUserPassword(c *gin.Context) {
	User := User{}
	if err := c.ShouldBindJSON(&User); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	if User.Name == "" || User.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input",
		})
	}
	err := model.UpdateUserPassword(User.Name, User.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "update user password sucessfully",
		})
	}
}

func DeleteUser(c *gin.Context) {
	User := User{}
	if err := c.ShouldBindJSON(&User); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	if User.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input",
		})
	}
	err := model.DeleteUser(User.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "delete user sucessfully",
	})
}

func UserLogin(c *gin.Context) {
	User := User{}
	if err := c.ShouldBindJSON(&User); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	userinfo, err := model.QueryUserInfo(User.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	userPasswordHash := utils.HashPassword(User.Password, userinfo.Token)
	if userinfo.Password == userPasswordHash {
		token, err := middleware.GenerateJWT(userinfo.Id, userinfo.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": token,
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "input wrong password",
		})
	}
}
