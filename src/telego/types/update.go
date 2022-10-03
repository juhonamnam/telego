package types

type Update struct {
	UpdateId          int          `json:"update_id"`
	Message           *Message     `json:"message"`
	EditedMessage     *Message     `json:"edited_message"`
	ChannelPost       *Message     `json:"channel_post"`
	EditedChannelPost *Message     `json:"edited_channel_post"`
	InlineQuery       *InlineQuery `json:"inline_query"`
	// chosen_inline_result
	CallbackQuery *CallbackQuery `json:"callback_query"`
	// shipping_query
	// pre_checkout_query
	// poll
	Poll *Poll `json:"poll"`
	// poll_answer
	// my_chat_member
	// chat_member
	// chat_join_request
}
