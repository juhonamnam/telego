package types

type LoginUrl struct {
	Url                string  `json:"url"`
	ForwardText        *string `json:"forward_text"`
	BotUsername        *string `json:"bot_username"`
	RequestWriteAccess *string `json:"request_write_access"`
}
