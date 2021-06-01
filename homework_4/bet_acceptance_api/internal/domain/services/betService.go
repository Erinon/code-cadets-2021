package services

import (
	"context"

	"code-cadets-2021/homework_4/bets_api/internal/domain/models"
)

// BetService implements bet related functions.
type BetService struct {
	betRepository BetRepository
}

// NewBetService creates a new instance of BetService.
func NewBetService(betRepository BetRepository) *BetService {
	return &BetService{
		betRepository: betRepository,
	}
}

// GetBet sends event update message to the queues.
func (e *BetService) GetBet(ctx context.Context, betId string) (*models.Bet, error) {
	domainBet, exists, err := e.betRepository.GetBetByID(ctx, betId)

	if !exists {
		return nil, err
	}

	return &domainBet, err
}

// GetUserBets sends event update message to the queues.
func (e *BetService) GetUserBets(ctx context.Context, userId string) ([]models.Bet, error) {
	domainBets, _, err := e.betRepository.GetBetsByCustomerID(ctx, userId)

	return domainBets, err
}

// GetBetsWithStatus sends event update message to the queues.
func (e *BetService) GetBetsWithStatus(ctx context.Context, status string) ([]models.Bet, error) {
	domainBets, _, err := e.betRepository.GetBetsByStatus(ctx, status)

	return domainBets, err
}
