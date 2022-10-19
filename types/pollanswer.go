package types

type PollAnswer struct {
	PollId    string `json:"poll_id"`
	User      User   `json:"user"`
	OptionIds []int  `json:"option_ids"`
}
