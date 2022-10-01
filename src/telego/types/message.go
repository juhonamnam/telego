package types

type Message struct {
	MessageId            int           `json:"message_id"`
	From                 *User         `json:"from"`
	SenderChat           *Chat         `json:"sender_chat"`
	Date                 int           `json:"date"`
	Chat                 Chat          `json:"chat"`
	ForwardFrom          *User         `json:"forward_from"`
	ForwardFromChat      *Chat         `json:"forward_from_chat"`
	ForwardFromMessageId *int          `json:"forward_from_message_id"`
	ForwardSignature     *string       `json:"forward_signature"`
	ForwardSenderName    *string       `json:"forward_sender_name"`
	ForwardDate          *int          `json:"forward_date"`
	IsAutomaticForward   *optionalTrue `json:"is_automatic_forward"`
	ReplyToMessage       *Message      `json:"reply_to_message"`
	ViaBot               *User         `json:"via_bot"`
	EditDate             *int          `json:"edit_date"`
	HasProtectedContent  *optionalTrue `json:"has_protected_content"`
	MediaGroupId         *string       `json:"media_group_id"`
	AuthorSignature      *string       `json:"author_signature"`
	Text                 *string       `json:"text"`
	// entities
	// animation
	// audio
	// document
	// photo
	// sticker
	// video
	// video_note
	// voice
	Caption *string `json:"caption"`
	// caption_entities
	// contact
	// dice
	// game
	// poll
	// venue
	// location
	NewChatMembers *[]*User `json:"new_chat_members"`
	LeftChatMember *User    `json:"left_chat_member"`
	NewChatTitle   *string  `json:"new_chat_title"`
	// new_chat_photo
	DeleteChatPhoto       *optionalTrue `json:"delete_chat_photo"`
	GroupChatCreated      *optionalTrue `json:"group_chat_created"`
	SupergroupChatCreated *optionalTrue `json:"supergroup_chat_created"`
	ChannelChatCreated    *optionalTrue `json:"channel_chat_created"`
	// message_auto_delete_timer_changed
	MigrateToChatId   *int     `json:"migrate_to_chat_id"`
	MigrateFromChatId *int     `json:"migrate_from_chat_id"`
	PinnedMessage     *Message `json:"pinned_message"`
	// invoice
	// successful_payment
	ConnectedWebsite *string `json:"connected_website"`
	// passport_data
	// proximity_alert_triggered
	// video_chat_scheduled
	// video_chat_started
	// video_chat_ended
	// video_chat_participants_invited
	// web_app_data
	// reply_markup
}
