package telego

import "net/http"

type telegoInterface interface {
	SetUpdateHandler(func(updateContext *Context))
	SetTimeout(timeout uint16)
	Start()
	Request(endpoint string, data any) (*[]byte, error)
}

type telegoStruct struct {
	apiKey        string
	logger        Logger
	offset        int
	botInfo       *getMeResult
	updateHandler func(updateContext *Context)
	timeout       *uint16
	httpClient    *http.Client
}

func (telego *telegoStruct) SetUpdateHandler(updateHandler func(updateContext *Context)) {
	telego.updateHandler = updateHandler
}

func (telego *telegoStruct) SetTimeout(timeout uint16) {
	telego.timeout = &timeout
}

func Initialize(apiKey string) telegoInterface {
	return &telegoStruct{
		apiKey: apiKey,
	}
}
