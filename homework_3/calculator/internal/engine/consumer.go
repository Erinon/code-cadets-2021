package engine

import (
	"context"

	rabbitmqmodels "code-cadets-2021/homework_3/calculator/internal/infrastructure/rabbitmq/models"
)

type Consumer interface {
	ConsumeBets(ctx context.Context) (<-chan rabbitmqmodels.Bet, error)
	ConsumeEventUpdates(ctx context.Context) (<-chan rabbitmqmodels.EventUpdate, error)
}
