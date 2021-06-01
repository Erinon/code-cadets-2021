package handler

import (
	"context"

	domainmodels "code-cadets-2021/homework_3/calculator/internal/domain/models"
)

type BetRepository interface {
	InsertBet(ctx context.Context, bet domainmodels.Bet) error
	GetBetsBySelectionID(ctx context.Context, id string) ([]domainmodels.Bet, error)
}
