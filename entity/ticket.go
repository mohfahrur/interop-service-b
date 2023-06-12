package entity

type SendTelegramRequest struct {
	User string `json:"user"`
	Item string `json:"item"`
}
