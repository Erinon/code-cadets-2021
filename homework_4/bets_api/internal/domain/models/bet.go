package models

// Bet bet model.
type Bet struct {
	Id      			 string
	CustomerId           string
	Status               string
	SelectionId          string
	SelectionCoefficient int
	Payment              int
	Payout               int
}
