package message

import (
	"github.com/juhonamnam/telego/src/telego"
)

type sendMessage = struct {
	ChatId      int          `json:"chat_id"`
	Text        string       `json:"text"`
	ReplyMarkup *replyMarkup `json:"reply_markup"`
}

type replyMarkup struct {
	InlineKeyboard *[][]inlineKeyboard `json:"inline_keyboard"`
}

type inlineKeyboard struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data"`
}

func HandleMessage(ctx *telego.Context) {
	if ctx.Update.Message.Text != nil {
		ctx.Request("sendMessage", &sendMessage{
			ChatId: ctx.Update.Message.Chat.Id,
			Text:   "Message: " + *ctx.Update.Message.Text,
			ReplyMarkup: &replyMarkup{
				InlineKeyboard: &[][]inlineKeyboard{
					{
						{
							Text:         "Callback",
							CallbackData: "callback_data",
						},
					},
				},
			},
		})
	}
}
