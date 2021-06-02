package consumer

import (
	"context"

	rabbitmqmodels "code-cadets-2021/homework_3/calculator/internal/infrastructure/rabbitmq/models"
)

type EventUpdateConsumer interface {
	Consume(ctx context.Context) (<-chan rabbitmqmodels.EventUpdate, error)
}
