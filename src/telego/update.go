package telego

import (
	"encoding/json"
	"errors"

	"github.com/juhonamnam/telego/src/telego/types"
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
	UpdateId          int            `json:"update_id"`
	Message           *types.Message `json:"message"`
	EditedMessage     *types.Message `json:"edited_message"`
	ChannelPost       *types.Message `json:"channel_post"`
	EditedChannelPost *types.Message `json:"edited_channel_post"`
	// inline_query
	// chosen_inline_result
	CallbackQuery *types.CallbackQuery `json:"callback_query"`
	// shipping_query
	// pre_checkout_query
	// poll
	// poll_answer
	// my_chat_member
	// chat_member
	// chat_join_request
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

type Context struct {
	Update *update
	telego telegoInterface
}

func (updateContext *Context) Request(endpoint string, data any) {
	updateContext.telego.Request(endpoint, data)
}

func (telego *telegoStruct) updateLoop() {
	telego.logger.Info("Update loop started")

	for {
		updates, err := telego.getUpdates(600, telego.offset)

		if err != nil {
			continue
		}

		for _, update := range *updates {
			nextOffset := update.UpdateId + 1
			if telego.offset < nextOffset {
				telego.offset = nextOffset
			}
			telego.updateHandler(&Context{
				telego: telego,
				Update: update,
			})
		}
	}
}
