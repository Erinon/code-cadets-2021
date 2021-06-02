package rabbitmq

import (
	"context"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"

	"github.com/pkg/errors"

	"code-cadets-2021/homework_3/calculator/internal/infrastructure/rabbitmq/models"
)

// BetCalculatedPublisher publishes calculated bets to the desired RabbitMQ queue.
type BetCalculatedPublisher struct {
	channel Channel
	config  PublisherConfig
}

// NewBetCalculatedPublisher creates and returns a new BetCalculatedPublisher.
func NewBetCalculatedPublisher(channel Channel, config PublisherConfig) (*BetCalculatedPublisher, error) {
	_, err := channel.QueueDeclare(
		config.Queue,
		config.DeclareDurable,
		config.DeclareAutoDelete,
		config.DeclareExclusive,
		config.DeclareNoWait,
		config.DeclareArgs,
	)
	if err != nil {
		return nil, errors.Wrap(err, "bet calculated publisher initialization failed")
	}

	return &BetCalculatedPublisher{
		channel: channel,
		config:  config,
	}, nil
}

// Publish publishes messages until the context is cancelled.
func (c *BetCalculatedPublisher) Publish(ctx context.Context, betsCalculated <-chan models.BetCalculated) {
	go func() {
		for betCalculated := range betsCalculated {
			select {
			case <-ctx.Done():
				return
			default:
				betJson, err := json.Marshal(&betCalculated)
				if err != nil {
					log.Println("Failed to marshal the following bet:", betCalculated)
					continue
				}

				err = c.channel.Publish(
					c.config.Exchange,
					c.config.Queue,
					c.config.Mandatory,
					c.config.Immediate,
					amqp.Publishing{
						ContentType: "text/plain",
						Body:        betJson,
					},
				)
				if err != nil {
					log.Println("Failed to publish the following bet:", betJson)
				}
			}
		}
	}()
}
