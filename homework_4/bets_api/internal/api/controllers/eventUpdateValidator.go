package controllers

// BetStatusValidator validates bet status requests.
type BetStatusValidator interface {
	BetStatusIsValid(status string) bool
}
