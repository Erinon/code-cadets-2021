package sqlite

import (
	domainmodels "code-cadets-2021/homework_3/calculator/internal/domain/models"
	storagemodels "code-cadets-2021/homework_3/calculator/internal/infrastructure/sqlite/models"
)

type BetMapper interface {
	MapDomainBetToStorageBet(domainBet domainmodels.Bet) storagemodels.Bet
	MapStorageBetToDomainBet(storageBet storagemodels.Bet) domainmodels.Bet
}
