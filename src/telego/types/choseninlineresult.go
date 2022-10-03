package types

type ChosenInlineResult struct {
	ResultId        string    `json:"result_id"`
	From            User      `json:"from"`
	Location        *Location `json:"location"`
	InlineMessageId *string   `json:"inline_message_id"`
	Query           string    `json:"query"`
}
