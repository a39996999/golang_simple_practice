package room

import (
	"chatroom/middleware"
	"chatroom/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Room struct {
	Room_id     int `json:"Room_id"`
	Owner_id    int
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

func CreateRoom(c *gin.Context) {
	room := Room{}
	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
		})
		return
	}
	token := c.GetHeader("token")
	tokenClaimes, _ := middleware.ParseToken(token)
	if room.Name == "" || room.Description == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
		})
		return
	}
	room.Owner_id = int(tokenClaimes["user_id"].(float64))
	err := model.CreateRoom(room.Name, room.Owner_id, room.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Create sucessfully",
		})
	}
}

func DeleteRoom(c *gin.Context) {
	room := Room{}
	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
		})
	}
	token := c.GetHeader("token")
	tokenClaimes, err := middleware.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	roomOwnerId, err := model.GetOwnerId(room.Room_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	reqOwnerId := int(tokenClaimes["user_id"].(float64))
	if roomOwnerId != reqOwnerId {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid operate",
		})
		return
	}
	err = model.DeleteRoom(room.Room_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Delete room sucessfully",
		})
	}
}

func GetAllRoom(c *gin.Context) {
	roomlist, err := model.GetRoomList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": roomlist,
		})
	}
}
