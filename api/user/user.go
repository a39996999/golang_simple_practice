package user

import (
	"chatroom/middleware"
	"chatroom/model"
	"chatroom/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	User_id  int
	Name     string `json:"Username"`
	Password string `json:"Password"`
	Email    string `json:"Email"`
}

func CreateUser(c *gin.Context) {
	User := User{}
	if err := c.ShouldBindJSON(&User); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if User.Name == "" || User.Password == "" || User.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid input",
		})
		return
	}
	checkEmailValidate, emailError := utils.VerifyEmailFormat(User.Email)
	if emailError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": emailError.Error(),
		})
	} else if checkEmailValidate == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Input wrong email format",
		})
	} else {
		user, err := model.CreateUser(User.Name, User.Password, User.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		} else if user.Email != "" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Email is exist",
			})
		} else if user.Name != "" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Name is exist",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "create user sucessfully",
			})
		}
	}
}
func CheckUserExist(c *gin.Context) {
	username := c.Param("username")
	user, err := model.QueryUserInfo(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	if user.Name != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Username is exist",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Username can used",
		})
	}
}

func UpdateUserPassword(c *gin.Context) {
	User := User{}
	if err := c.ShouldBindJSON(&User); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	if User.Name == "" || User.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid input",
		})
	}
	err := model.UpdateUserPassword(User.Name, User.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "update user password sucessfully",
		})
	}
}

func DeleteUser(c *gin.Context) {
	User := User{}
	if err := c.ShouldBindJSON(&User); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	if User.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid input",
		})
	}
	err := model.DeleteUser(User.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "delete user sucessfully",
	})
}

func UserLogin(c *gin.Context) {
	User := User{}
	if err := c.ShouldBindJSON(&User); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	userinfo, err := model.QueryUserInfo(User.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	if userinfo.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User is not exist",
		})
		return
	}

	userPasswordHash := utils.HashPassword(User.Password, userinfo.Token)
	if userinfo.Password == userPasswordHash {
		if userinfo.IsVerify == false {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Email is not verified",
			})
			return
		}
		token, err := middleware.GenerateJWT(userinfo.Id, userinfo.Name, userinfo.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": token,
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "input wrong password",
		})
	}
}

func UserStatus(c *gin.Context) {
	token := c.GetHeader("token")
	tokenClaimes, err := middleware.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
	}
	user := User{}
	user.User_id = int(tokenClaimes["user_id"].(float64))
	user.Name = tokenClaimes["username"].(string)
	user.Email = tokenClaimes["email"].(string)
	c.JSON(http.StatusOK, gin.H{
		"message": gin.H{
			"user_id":  user.User_id,
			"username": user.Name,
			"email":    user.Email,
		},
	})
}
