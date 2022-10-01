package callback

import "github.com/juhonamnam/telego/src/telego"

type editMessageText = struct {
	ChatId    int    `json:"chat_id"`
	MessageId int    `json:"message_id"`
	Text      string `json:"text"`
}

func HandleCallback(ctx *telego.Context) {
	if ctx.Update.CallbackQuery.Data != nil {
		ctx.Request("editMessageText", &editMessageText{
			ChatId:    ctx.Update.CallbackQuery.Message.Chat.Id,
			MessageId: ctx.Update.CallbackQuery.Message.MessageId,
			Text:      "Callback: " + *ctx.Update.CallbackQuery.Data,
		})
	}
}
