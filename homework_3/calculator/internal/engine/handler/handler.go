package handler

import (
	"context"
	"log"

	domainmodels "code-cadets-2021/homework_3/calculator/internal/domain/models"
	rabbitmqmodels "code-cadets-2021/homework_3/calculator/internal/infrastructure/rabbitmq/models"
)

// Handler handles bets received and bets calculated.
type Handler struct {
	betRepository BetRepository
}

// New creates and returns a new Handler.
func New(betRepository BetRepository) *Handler {
	return &Handler{
		betRepository: betRepository,
	}
}

// HandleBets handles bets.
func (h *Handler) HandleBets(
	ctx context.Context,
	bets <-chan rabbitmqmodels.Bet,
) {
	go func() {
		for bet := range bets {
			log.Println("Processing bet, betId:", bet.Id)

			// Calculate the domain bet based on the incoming bet.
			domainBet := domainmodels.Bet{
				Id:                   bet.Id,
				SelectionId:          bet.SelectionId,
				SelectionCoefficient: bet.SelectionCoefficient,
				Payment:              bet.Payment,
			}

			// Insert the domain bet into the repository.
			err := h.betRepository.InsertBet(ctx, domainBet)
			if err != nil {
				log.Println("Failed to insert bet, error: ", err)
				continue
			}
		}
	}()
}

// HandleEventUpdates handles event updates.
func (h *Handler) HandleEventUpdates(
	ctx context.Context,
	eventUpdates <-chan rabbitmqmodels.EventUpdate,
) <-chan rabbitmqmodels.BetCalculated {
	resultingBets := make(chan rabbitmqmodels.BetCalculated)

	go func() {
		defer close(resultingBets)

		for eventUpdate := range eventUpdates {
			log.Println("Processing event update, betId:", eventUpdate.Id)

			// Fetch the domain bet.
			domainBets, err := h.betRepository.GetBetsBySelectionID(ctx, eventUpdate.Id)
			if err != nil {
				log.Println("Failed to fetch bets which should be updated, error: ", err)
				continue
			}
			if len(domainBets) == 0 {
				log.Println("Bets which should be updated do not exist, selectionId: ", eventUpdate.Id)
				continue
			}

			for _, domainBet := range domainBets {
				// Calculate the resulting bet, which should be published.
				resultingBet := rabbitmqmodels.BetCalculated{
					Id: domainBet.Id,
				}

				if eventUpdate.Outcome == "won" {
					resultingBet.Status = "won"
					resultingBet.Payout = domainBet.Payment * domainBet.SelectionCoefficient
				} else if eventUpdate.Outcome == "lost" {
					resultingBet.Status = "lost"
					resultingBet.Payout = 0
				} else {
					log.Println("Bets which should be updated do not exist, selectionId: ", eventUpdate.Id)
					break
				}

				select {
				case resultingBets <- resultingBet:
				case <-ctx.Done():
					return
				}
			}
		}
	}()

	return resultingBets
}
