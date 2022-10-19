package types

type ChatInviteLink struct {
	InviteLink              string  `json:"invite_link"`
	Creator                 User    `json:"creator"`
	CreatesJoinRequest      bool    `json:"creates_join_request"`
	IsPrimary               bool    `json:"is_primary"`
	IsRevoked               bool    `json:"is_revoked"`
	Name                    *string `json:"name"`
	ExpireDate              *int    `json:"expire_date"`
	MemberLimit             *int    `json:"member_limit"`
	PendingJoinRequestCount *int    `json:"pending_join_request_count"`
}
