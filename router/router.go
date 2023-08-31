package router

import (
	"chatroom/api/mail"
	"chatroom/api/room"
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
	userapiGruop := server.Group("/v1/user")
	{
		userapiGruop.POST("/create", user.CreateUser)
		userapiGruop.POST("/updatePassword", user.UpdateUserPassword)
		userapiGruop.POST("/delete", user.DeleteUser)
		userapiGruop.POST("/login", user.UserLogin)
		userapiGruop.POST("/sendmail", mail.SendMailToken)
		userapiGruop.GET("/verifymail/:token", mail.VerifyMailCode)
		userapiGruop.GET("/query/:username", user.CheckUserExist)
		userapiGruop.Use(middleware.JWTVerifyToken())
		{
			userapiGruop.GET("/useralive", user.UserStatus)
		}
	}
	roomapiGroup := server.Group("/v1/room")
	{
		roomapiGroup.Use(middleware.JWTVerifyToken())
		{
			roomapiGroup.POST("/create", room.CreateRoom)
			roomapiGroup.POST("/delete", room.DeleteRoom)
			roomapiGroup.GET("/getroomlist", room.GetAllRoom)
		}
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
