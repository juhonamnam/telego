package types

type Location struct {
	Longitude            float64  `json:"longitude"`
	Latitude             float64  `json:"latitude"`
	HorizontalAccuracy   *float64 `json:"horizontal_accuracy"`
	LivePeriod           *int     `json:"live_period"`
	Heading              *int     `json:"heading"`
	ProximityAlertRadius *int     `json:"proximity_alert_radius"`
}
