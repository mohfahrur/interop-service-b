package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	telegramD "github.com/mohfahrur/interop-service-b/domain/telegram"
	"github.com/mohfahrur/interop-service-b/entity"
	ticketUC "github.com/mohfahrur/interop-service-b/usecase/ticket"
)

func main() {
	token := os.Getenv("token")
	chatID := os.Getenv("chatid")
	telegramDomain := telegramD.NewTelegramDomain(token, chatID)
	ticketUsecase := ticketUC.NewTicketUsecase(*telegramDomain)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong from service b",
		})
	})
	r.POST("/send-telegram", func(c *gin.Context) {

		var req entity.SendTelegramRequest
		err := c.BindJSON(&req)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "bad request",
			})
			return
		}
		err = ticketUsecase.SendTelegram(req)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "bad request",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})
	r.Run(":5001")
}
