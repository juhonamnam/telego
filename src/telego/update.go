package telego

import (
	"encoding/json"
	"errors"
)

type updateOption struct {
	Timeout int `json:"timeout"`
	Offset  int `json:"offset"`
}

type updateResponse struct {
	baseResponse
	Result *[]*update
}

type update struct {
	UpdateId int `json:"update_id"`
	Message  *message
}

type message struct {
	MessageId int `json:"message_id"`
	From      *From
}

type From struct {
	Id           int
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string
	LanguageCode string `json:"language_code"`
}

func (telego *telegoStruct) getUpdates(timeout int, offset int) (*[]*update, error) {
	endpoint := "getUpdates"
	res, err := telego.Request(endpoint, &updateOption{
		Timeout: timeout,
		Offset:  offset,
	})

	if err != nil {
		return nil, err
	}

	var updateRes updateResponse
	json.Unmarshal(*res, &updateRes)

	if !updateRes.Ok {
		telego.logger.Error(endpoint, updateRes.Description)
		return nil, errors.New(updateRes.Description)
	}

	return updateRes.Result, nil
}

type UpdateContext struct {
	Update *update
	telego action
}

func (updateContext *UpdateContext) SendMessage(chatId int, message string) {
	updateContext.telego.SendMessage(chatId, message)
}

func (updateContext *UpdateContext) Request(endpoint string, data any) {
	updateContext.Request(endpoint, data)
}

func (telego *telegoStruct) updateLoop() {
	telego.logger.Info("Update loop started")

	for true {
		updates, err := telego.getUpdates(600, telego.offset)

		if err != nil {
			continue
		}

		for _, update := range *updates {
			nextOffset := update.UpdateId + 1
			if telego.offset < nextOffset {
				telego.offset = nextOffset
			}
			telego.updateHandler(&UpdateContext{
				telego: telego,
				Update: update,
			})
		}
	}
}
