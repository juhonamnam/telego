package types

type MessageEntity struct {
	Type          string  `json:"type"`
	Offset        int     `json:"offset"`
	Length        int     `json:"length"`
	Url           *string `json:"url"`
	User          *User   `json:"user"`
	Language      *string `json:"language"`
	CustomEmojiId *string `json:"custom_emoji_id"`
}
