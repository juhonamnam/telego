package telego

type telegoInterface interface {
	SetUpdateHandler(func(updateContext *Context))
	Start()
	Request(endpoint string, data any) (*[]byte, error)
}

type telegoStruct struct {
	apiKey        string
	logger        Logger
	offset        int
	botInfo       *getMeResult
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
