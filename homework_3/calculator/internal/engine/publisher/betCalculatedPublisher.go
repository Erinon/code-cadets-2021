package publisher

import (
	"context"

	rabbitmqmodels "code-cadets-2021/homework_3/calculator/internal/infrastructure/rabbitmq/models"
)

type BetCalculatedPublisher interface {
	Publish(ctx context.Context, betsCalculated <-chan rabbitmqmodels.BetCalculated)
}
