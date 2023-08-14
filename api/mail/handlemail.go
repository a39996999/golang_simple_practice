package mail

import (
	"chatroom/model"
	"chatroom/utils"
	_ "chatroom/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmailRequest struct {
	Email string `json:"Email"`
}

func SendMailToken(c *gin.Context) {
	email := EmailRequest{}
	if err := c.ShouldBindJSON(&email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  10001,
			"error": err.Error(),
		})
		return
	}
	if email.Email == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":  10001,
			"error": "invalid email",
		})
		return
	}
	mailToken, err := utils.GenerateToken()
	verifyUrl := "http://127.0.0.1:8080/v1/user/verifymail/" + mailToken
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  10001,
			"error": err.Error(),
		})
		return
	}
	err = model.RecordSendMail(email.Email, mailToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  10001,
			"error": err.Error(),
		})
		return
	}
	message := []byte("To: " + email.Email + "\r\n" +
		"Subject: Verify mail\r\n" + "\r\n" +
		verifyUrl + "\r\n")
	err = utils.SendMail(email.Email, message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  10001,
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "send mail sucessfully",
		})
	}
}

func VerifyMailCode(c *gin.Context) {
	token := c.Param("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  10001,
			"error": "token invalid",
		})
		return
	}
	err := model.VerifyMail(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  10001,
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "verify mail sucessfully",
		})
	}
}
