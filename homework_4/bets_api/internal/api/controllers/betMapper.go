package controllers

import (
	dtomodels "code-cadets-2021/homework_4/bets_api/internal/api/controllers/models"
	domainmodels "code-cadets-2021/homework_4/bets_api/internal/domain/models"
)

type BetMapper interface {
	MapDomainBetToBetDto(domainBet domainmodels.Bet) dtomodels.BetResponseDto
}
