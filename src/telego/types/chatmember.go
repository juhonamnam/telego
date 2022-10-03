package types

type ChatMember struct {
	Status                string  `json:"status"`
	User                  User    `json:"user"`
	IsAnonymous           *bool   `json:"is_anonymous"`
	CustomTitle           *string `json:"custom_title"`
	CanBeEdited           *bool   `json:"can_be_edited"`
	CanManageChat         *bool   `json:"can_manage_chat"`
	CanDeleteMessages     *bool   `json:"can_delete_messages"`
	CanManageVideoChats   *bool   `json:"can_manage_video_chats"`
	CanRestrictMembers    *bool   `json:"can_restrict_members"`
	CanPromoteMembers     *bool   `json:"can_promote_members"`
	CanChangeInfo         *bool   `json:"can_change_info"`
	CanInviteUsers        *bool   `json:"can_invite_users"`
	CanPostMessages       *bool   `json:"can_post_messages"`
	CanEditMessages       *bool   `json:"can_edit_messages"`
	CanPinMessages        *bool   `json:"can_pin_messages"`
	IsMember              *bool   `json:"is_member"`
	CanSendMessages       *bool   `json:"can_send_messages"`
	CanSendMediaMessages  *bool   `json:"can_send_media_messages"`
	CanSendPolls          *bool   `json:"can_send_polls"`
	CanSendOtherMessages  *bool   `json:"can_send_other_messages"`
	CanAddWebPagePreviews *bool   `json:"can_add_web_page_previews"`
	UntilDate             *int    `json:"until_date"`
}

type ChatMemberOwner struct {
	Status      string  `json:"status"`
	User        User    `json:"user"`
	IsAnonymous *bool   `json:"is_anonymous"`
	CustomTitle *string `json:"custom_title"`
}

type ChatMemberAdministrator struct {
	Status              string  `json:"status"`
	User                User    `json:"user"`
	CanBeEdited         *bool   `json:"can_be_edited"`
	IsAnonymous         *bool   `json:"is_anonymous"`
	CanManageChat       *bool   `json:"can_manage_chat"`
	CanDeleteMessages   *bool   `json:"can_delete_messages"`
	CanManageVideoChats *bool   `json:"can_manage_video_chats"`
	CanRestrictMembers  *bool   `json:"can_restrict_members"`
	CanPromoteMembers   *bool   `json:"can_promote_members"`
	CanChangeInfo       *bool   `json:"can_change_info"`
	CanInviteUsers      *bool   `json:"can_invite_users"`
	CanPostMessages     *bool   `json:"can_post_messages"`
	CanEditMessages     *bool   `json:"can_edit_messages"`
	CanPinMessages      *bool   `json:"can_pin_messages"`
	CustomTitle         *string `json:"custom_title"`
}

type ChatMemberMember struct {
	Status string `json:"status"`
	User   User   `json:"user"`
}

type ChatMemberRestricted struct {
	Status                string `json:"status"`
	User                  User   `json:"user"`
	IsMember              *bool  `json:"is_member"`
	CanChangeInfo         *bool  `json:"can_change_info"`
	CanInviteUsers        *bool  `json:"can_invite_users"`
	CanPinMessages        *bool  `json:"can_pin_messages"`
	CanSendMessages       *bool  `json:"can_send_messages"`
	CanSendMediaMessages  *bool  `json:"can_send_media_messages"`
	CanSendPolls          *bool  `json:"can_send_polls"`
	CanSendOtherMessages  *bool  `json:"can_send_other_messages"`
	CanAddWebPagePreviews *bool  `json:"can_add_web_page_previews"`
	UntilDate             *int   `json:"until_date"`
}

type ChatMemberLeft struct {
	Status string `json:"status"`
	User   User   `json:"user"`
}

type ChatMemberBanned struct {
	Status    string `json:"status"`
	User      User   `json:"user"`
	UntilDate *int   `json:"until_date"`
}
