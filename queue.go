package workqueue

import "context"

// WIP thinking about how to encode the problem.

type Work interface {
}

type Queue interface {
	NextWork(context.Context) (Work, error)
}

type QueueBackend interface {
	ReserveWork(context.Context, Work) error
	ScheduleRetry(context.Context, Work) error
	MarkComplete(context.Context, Work) error
}
