package main

import (
	"chatroom/model"
	"chatroom/router"
)

func init() {
	model.Init()
}
func main() {
	server := router.InitRouter()
	server.Run(":8080")
}
