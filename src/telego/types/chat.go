package types

type Chat struct {
	Id        int     `json:"id"`
	Type      string  `json:"type"`
	Title     *string `json:"title"`
	Username  *string `json:"username"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	// photo
	Bio                               *string
	HasPrivateForwards                *optionalTrue `json:"has_private_forwards"`
	HasRestrictedVoiceAndVideoMessage *optionalTrue `json:"has_restricted_voice_and_video_messages"`
	JoinToSendMessage                 *optionalTrue `json:"join_to_send_messages"`
	JoinByRequest                     *optionalTrue `json:"join_by_request"`
	Description                       *string       `json:"description"`
	InviteLink                        *string       `json:"invite_link"`
	PinnedMessage                     *Message      `json:"pinned_message"`
	// permissions
	SlowModeDelay         *int          `json:"slow_mode_delay"`
	MessageAutoDeleteTime *int          `json:"message_auto_delete_time"`
	HasProtectedContent   *optionalTrue `json:"has_protected_content"`
	StickerSetName        *string       `json:"sticker_set_name"`
	CanSetStickerSet      *optionalTrue `json:"can_set_sticker_set"`
	LinkedChatId          *int          `json:"linked_chat_id"`
	// location
}
