package services

import (
	"code-cadets-2021/lecture_2/05_offerfeed/internal/domain/models"
	"context"
)

type FeedProcessorService struct {
	feed Feed
	queue Queue
}

func NewFeedProcessorService(feed Feed, queue Queue) *FeedProcessorService {
	// it should receive "Feed" & "Queue" interfaces through constructor
	return &FeedProcessorService{feed: feed, queue: queue}
}

func (f *FeedProcessorService) Start(ctx context.Context) error {
	// initially:
	// - get updates channel from feed interface
	// - get source channel from queue interface
	updates := f.feed.GetUpdates()
	source :=  f.queue.GetSource()

	// repeatedly:
	// - range over updates channel
	// - multiply each odd with 2
	// - send it to source channel
	for update := range updates {
		update.Coefficient *= 2
		source <- update
	}

	// finally:
	// - when updates channel is closed, exit
	// - when exiting, close source channel
	close(source)

	return nil
}

// Feed define feed interface here
type Feed interface {
	GetUpdates() chan models.Odd
}

// Queue define queue interface here
type Queue interface {
	GetSource() chan models.Odd
}
