package router

import (
	"chatroom/api/mail"
	"chatroom/api/testing"
	"chatroom/api/user"
	"chatroom/middleware"
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	server := gin.Default()
	server.Use(static.Serve("/static", static.LocalFile("frontend/static", false)))
	server.LoadHTMLGlob("frontend/template/*")
	apiGroup := server.Group("/v1/user")
	{
		apiGroup.POST("/create", user.CreateUser)
		apiGroup.POST("/updatePassword", user.UpdateUserPassword)
		apiGroup.POST("/delete", user.DeleteUser)
		apiGroup.POST("/login", user.UserLogin)
		apiGroup.POST("/sendmail", mail.SendMailToken)
		apiGroup.GET("/verifymail/:token", mail.VerifyMailCode)
		apiGroup.GET("/query/:username", user.CheckUserExist)
	}
	apiTestGroup := server.Group("test")
	{
		apiTestGroup.GET("ping", middleware.JWTVerifyToken(), testing.Ping)
	}
	frontGroup := server.Group("home")
	{
		frontGroup.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{})
		})
	}
	return server
}
