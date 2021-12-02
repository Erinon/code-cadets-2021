package models

// BetResponseDto bet response dto model.
type BetResponseDto struct {
	Id      			 string `json:"id"`
	Status               string `json:"status"`
	SelectionId          string `json:"selection_id"`
	SelectionCoefficient int    `json:"selection_coefficient"`
	Payment              int    `json:"payment"`
	Payout               int    `json:"payout"`
}
