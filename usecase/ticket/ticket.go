package ticket

import (
	"fmt"
	"log"

	telegramD "github.com/mohfahrur/interop-service-b/domain/telegram"
	"github.com/mohfahrur/interop-service-b/entity"
)

type TicketAgent interface {
	SendEmail(user string, item string) (err error)
}

type TicketUsecase struct {
	telegramDomain telegramD.TelegramDomain
}

func NewTicketUsecase(
	telegramD telegramD.TelegramDomain) *TicketUsecase {

	return &TicketUsecase{
		telegramDomain: telegramD}
}

func (uc *TicketUsecase) SendTelegram(req entity.SendTelegramRequest) (err error) {
	msg := fmt.Sprintf("NOTIFICATION: Pembelian tiket film %s, oleh %s", req.Item, req.User)

	err = uc.telegramDomain.SendMessage(msg)
	if err != nil {
		log.Println(err)
	}
	return
}
