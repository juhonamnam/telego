package types

type ProximityAlertTriggered struct {
	Traveler User `json:"traveler"`
	Watcher  User `json:"watcher"`
	Distance int  `json:"distance"`
}
