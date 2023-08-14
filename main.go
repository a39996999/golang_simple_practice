package main

import (
	"chatroom/model"
	"chatroom/router"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	model.InitDB()
}
func main() {
	server := router.InitRouter()
	server.Run(os.Getenv("server_host") + ":" + os.Getenv("server_port"))
}
