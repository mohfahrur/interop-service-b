package telegram

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type TelegramAgent interface {
	SendMessage(messageBody string) error
}

type TelegramDomain struct {
	Token  string
	ChatID int64
}

func NewTelegramDomain(token, chatID string) *TelegramDomain {
	i, err := strconv.ParseInt(chatID, 10, 64)
	if err != nil {
		panic(err)
	}
	return &TelegramDomain{
		Token:  token,
		ChatID: i,
	}
}

func (d *TelegramDomain) SendMessage(messageBody string) (err error) {

	bot, err := tgbotapi.NewBotAPI(d.Token)
	if err != nil {
		log.Fatal(err)
	}

	msg := tgbotapi.NewMessage(d.ChatID, messageBody)

	_, err = bot.Send(msg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Message sent successfully!")
	return
}
