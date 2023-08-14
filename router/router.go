package router

import (
	"chatroom/api/mail"
	"chatroom/api/user"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	server := gin.Default()
	apiGroup := server.Group("/v1/user")
	{
		apiGroup.POST("/create", user.CreateUser)
		apiGroup.POST("/updatePassword", user.UpdateUserPassword)
		apiGroup.POST("/delete", user.DeleteUser)
		apiGroup.POST("/login", user.UserLogin)
		apiGroup.POST("/sendmail", mail.SendMailToken)
		apiGroup.GET("/verifymail/:token", mail.VerifyMailCode)
	}
	return server
}
