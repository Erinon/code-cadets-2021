package mappers

import (
	dtomodels "code-cadets-2021/homework_4/bets_api/internal/api/controllers/models"
	domainmodels "code-cadets-2021/homework_4/bets_api/internal/domain/models"
)

// BetMapper maps bet dtos to domain bets and vice versa.
type BetMapper struct {
}

// NewBetMapper creates and returns a new BetMapper.
func NewBetMapper() *BetMapper {
	return &BetMapper{}
}

// MapDomainBetToBetDto maps the given domain bet into bet dto.
func (m *BetMapper) MapDomainBetToBetDto(domainBet domainmodels.Bet) dtomodels.BetResponseDto {
	return dtomodels.BetResponseDto{
		Id:                   domainBet.Id,
		Status:               domainBet.Status,
		SelectionId:          domainBet.SelectionId,
		SelectionCoefficient: domainBet.SelectionCoefficient,
		Payment:              domainBet.Payment,
		Payout:               domainBet.Payout,
	}
}
