package types

type Sticker struct {
	FileId           string        `json:"file_id"`
	FileUniqueId     string        `json:"file_unique_id"`
	Type             string        `json:"type"`
	Width            int           `json:"width"`
	Height           int           `json:"height"`
	IsAnimated       bool          `json:"is_animated"`
	IsVideo          bool          `json:"is_video"`
	Thumb            *PhotoSize    `json:"thumb"`
	Emoji            *string       `json:"emoji"`
	SetName          *string       `json:"set_name"`
	PremiumAnimation *File         `json:"premium_animation"`
	MaskPosition     *MaskPosition `json:"mask_position"`
	CustomEmojiId    *string       `json:"custom_emoji_id"`
	FileSize         *int          `json:"file_size"`
}
