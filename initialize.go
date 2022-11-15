package telego

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/juhonamnam/telego/types"
)

type getMeResponse struct {
	baseResponse
	Result *getMeResult
}

type getMeResult struct {
	types.User
	CanJoinGroups           *bool `json:"can_join_groups"`
	CanReadAllGroupMessages *bool `json:"can_read_all_group_messages"`
	SupportsInlineQueries   *bool `json:"supports_inline_queries"`
}

func (telego *telegoStruct) getMe() {
	telego.logger.Info("Getting Bot Information...")

	response, err := telego.Request("getMe", nil)

	if err != nil {
		panic(err.Error())
	}

	var getMeRes getMeResponse
	json.Unmarshal(*response, &getMeRes)

	if !getMeRes.Ok {
		panic(getMeRes.Description)
	}

	telego.botInfo = getMeRes.Result
}

func (telego *telegoStruct) getInitialOffset() {
	telego.logger.Info("Setting initial offset...")

	updates, err := telego.getUpdates(0, 0)

	if err != nil {
		panic(err.Error())
	}

	for _, update := range *updates {
		nextOffset := update.UpdateId + 1
		if update.UpdateId < nextOffset {
			telego.offset = nextOffset
		}
	}
}

func (telego *telegoStruct) Start() {
	if telego.logger == nil {
		telego.logger = getInitalLogger()
	}
	if telego.updateHandler == nil {
		telego.updateHandler = func(updateContext *Context) {}
	}
	if telego.timeout == nil {
		var timeout uint16 = 600
		telego.timeout = &timeout
	}
	if len(telego.apiKey) == 0 {
		panic("Telegram bot API not found")
	}
	telego.httpClient = &http.Client{Timeout: time.Duration(*telego.timeout) * time.Second}
	telego.getMe()
	telego.getInitialOffset()
	telego.logger.Info("Ready")
	telego.updateLoop()
}
