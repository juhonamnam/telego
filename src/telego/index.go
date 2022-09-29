package telego

type telegoInterface interface {
	initializer
	action
}

type initializer interface {
	SetUpdateHandler(func(updateContext *UpdateContext))
	Start()
}
type action interface {
	SendMessage(chatId int, message string)
	Request(endpoint string, data any) (*[]byte, error)
}

type telegoStruct struct {
	apiKey        string
	logger        Logger
	offset        int
	botInfo       *botInfo
	updateHandler func(updateContext *UpdateContext)
}

type botInfo struct {
	Username string
}

func (telego *telegoStruct) SetUpdateHandler(updateHandler func(updateContext *UpdateContext)) {
	telego.updateHandler = updateHandler
}

func Default(apiKey string) telegoInterface {
	return &telegoStruct{
		apiKey:        apiKey,
		logger:        getInitalLogger(),
		updateHandler: func(updateContext *UpdateContext) {},
	}
}
