package router

import (
	"chatroom/api/user"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	server := gin.Default()
	apiGroup := server.Group("/v1/user")
	{
		apiGroup.POST("/create", user.CreateUser)
		apiGroup.POST("/updatepassword", user.UpdateUserPassword)
		apiGroup.POST("/delete", user.DeleteUser)
		apiGroup.POST("/login", user.UserLogin)
	}
	return server
}
