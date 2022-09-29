package message

import "github.com/juhonamnam/telego/src/telego"

type sendMessage = struct {
	ChatId int    `json:"chat_id"`
	Text   string `json:"text"`
}

func HandleMessage(ctx *telego.Context) {
	message := sendMessage{
		ChatId: ctx.Update.Message.From.Id,
		Text:   "Hello",
	}
	ctx.Request("sendMessage", message)
}
