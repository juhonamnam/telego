package telego

import "encoding/json"

type sendMessage = struct {
	ChatId int    `json:"chat_id"`
	Text   string `json:"text"`
}

func (telego *telegoStruct) SendMessage(chatId int, text string) {
	endpoint := "sendMessage"
	res, err := telego.Request(endpoint, &sendMessage{
		ChatId: chatId,
		Text:   text,
	})

	if err != nil {
		return
	}

	var sendMesRes baseResponse
	json.Unmarshal(*res, &sendMesRes)

	if !sendMesRes.Ok {
		telego.logger.Error(endpoint, sendMesRes.Description)
	}
}
