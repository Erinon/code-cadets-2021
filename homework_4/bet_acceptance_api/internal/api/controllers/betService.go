package controllers

import (
	"context"

	"code-cadets-2021/homework_4/bets_api/internal/domain/models"
)

// BetService implements bet related functions.
type BetService interface {
	GetBet(ctx context.Context, betId string) (*models.Bet, error)
	GetUserBets(ctx context.Context, userId string) ([]models.Bet, error)
	GetBetsWithStatus(ctx context.Context, status string) ([]models.Bet, error)
}
