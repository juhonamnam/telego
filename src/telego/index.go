package telego

import "github.com/juhonamnam/telego/src/telego/types"

type telegoInterface interface {
	SetUpdateHandler(func(updateContext *Context))
	Start()
	Request(endpoint string, data any) (*[]byte, error)
}

type telegoStruct struct {
	apiKey        string
	logger        Logger
	offset        int
	botInfo       *types.User
	updateHandler func(updateContext *Context)
}

func (telego *telegoStruct) SetUpdateHandler(updateHandler func(updateContext *Context)) {
	telego.updateHandler = updateHandler
}

func Default(apiKey string) telegoInterface {
	return &telegoStruct{
		apiKey:        apiKey,
		logger:        getInitalLogger(),
		updateHandler: func(updateContext *Context) {},
	}
}
