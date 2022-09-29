package telego

import (
	"encoding/json"
)

type getMeResponse struct {
	baseResponse
	Result *botInfo
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
	telego.getMe()
	telego.getInitialOffset()
	telego.logger.Info("Ready")
	telego.updateLoop()
}
