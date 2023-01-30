package telego

import (
	"net/http"
	"time"
)

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
	timeout       uint16
	httpClient    *http.Client
}

func (telego *telegoStruct) SetUpdateHandler(updateHandler func(updateContext *Context)) {
	telego.updateHandler = updateHandler
}

func (telego *telegoStruct) SetTimeout(timeout uint16) {
	telego.timeout = timeout
	telego.httpClient = &http.Client{Timeout: time.Duration(timeout) * time.Second}
}

func (telego *telegoStruct) SetLogger(logger Logger) {
	telego.logger = logger
}

func Initialize(apiKey string) telegoInterface {
	if len(apiKey) == 0 {
		panic("Telegram bot API not found")
	}
	return &telegoStruct{
		apiKey:     apiKey,
		logger:     getInitalLogger(),
		timeout:    600,
		httpClient: &http.Client{Timeout: time.Duration(600) * time.Second},
	}
}
