package types

type ChatJoinRequest struct {
	Chat       Chat            `json:"chat"`
	From       User            `json:"from"`
	Date       int             `json:"date"`
	Bio        *string         `json:"bio"`
	InviteLink *ChatInviteLink `json:"invite_link"`
}
