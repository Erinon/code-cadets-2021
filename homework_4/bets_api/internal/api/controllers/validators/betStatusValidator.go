package validators

const activeStatus = "active"
const wonStatus = "won"
const lostStatus = "lost"

// BetStatusValidator validates bet status requests.
type BetStatusValidator struct{}

// NewBetStatusValidator creates a new instance of BetStatusValidator.
func NewBetStatusValidator() *BetStatusValidator {
	return &BetStatusValidator{}
}

// BetStatusIsValid checks if status is valid.
// Status is `active`, `won`or `lost`
func (e *BetStatusValidator) BetStatusIsValid(status string) bool {
	return status == activeStatus || status == wonStatus || status == lostStatus
}
