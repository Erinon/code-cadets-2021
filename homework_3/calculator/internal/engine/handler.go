package engine

import (
	"context"

	rabbitmqmodels "code-cadets-2021/homework_3/calculator/internal/infrastructure/rabbitmq/models"
)

type Handler interface {
	HandleBets(ctx context.Context, bets <-chan rabbitmqmodels.Bet)
	HandleEventUpdates(ctx context.Context, eventUpdates <-chan rabbitmqmodels.EventUpdate) <-chan rabbitmqmodels.BetCalculated
}
