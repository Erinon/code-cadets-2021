package consumer

import (
	"context"

	rabbitmqmodels "code-cadets-2021/homework_3/calculator/internal/infrastructure/rabbitmq/models"
)

// Consumer offers methods for consuming from input queues.
type Consumer struct {
	betConsumer         BetConsumer
	eventUpdateConsumer EventUpdateConsumer
}

// New creates and returns a new Consumer.
func New(betConsumer BetConsumer, eventUpdateConsumer EventUpdateConsumer) *Consumer {
	return &Consumer{
		betConsumer:         betConsumer,
		eventUpdateConsumer: eventUpdateConsumer,
	}
}

// ConsumeBets consumes bets queue.
func (c *Consumer) ConsumeBets(ctx context.Context) (<-chan rabbitmqmodels.Bet, error) {
	return c.betConsumer.Consume(ctx)
}

// ConsumeEventUpdates consumes event updates queue.
func (c *Consumer) ConsumeEventUpdates(ctx context.Context) (<-chan rabbitmqmodels.EventUpdate, error) {
	return c.eventUpdateConsumer.Consume(ctx)
}
