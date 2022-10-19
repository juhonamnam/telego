package types

type User struct {
	Id                    int           `json:"id"`
	IsBot                 bool          `json:"is_bot"`
	FirstName             string        `json:"first_name"`
	LastName              *string       `json:"last_name"`
	Username              *string       `json:"username"`
	LanguageCode          *string       `json:"language_code"`
	IsPremium             *optionalTrue `json:"is_premium"`
	AddedToAttachmentMenu *optionalTrue `json:"added_to_attachment_menu"`
	/* Fields below are returned only in getMe request */
	// can_join_groups
	// can_read_all_group_messages
	// supports_inline_queries
}
