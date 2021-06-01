package mappers

import (
	domainmodels "code-cadets-2021/homework_4/bets_api/internal/domain/models"
	storagemodels "code-cadets-2021/homework_4/bets_api/internal/infrastructure/sqlite/models"
)

// BetMapper maps bet dtos to domain bets and vice versa.
type BetMapper struct {
}

// NewBetMapper creates and returns a new BetMapper.
func NewBetMapper() *BetMapper {
	return &BetMapper{}
}

// MapStorageBetToDomainBet maps the given storage bet into domain bet.
func (m *BetMapper) MapStorageBetToDomainBet(storageBet storagemodels.Bet) domainmodels.Bet {
	return domainmodels.Bet{
		Id:                   storageBet.Id,
		CustomerId: 		  storageBet.CustomerId,
		Status:               storageBet.Status,
		SelectionId:          storageBet.SelectionId,
		SelectionCoefficient: storageBet.SelectionCoefficient,
		Payment:              storageBet.Payment,
		Payout:               storageBet.Payout,
	}
}
