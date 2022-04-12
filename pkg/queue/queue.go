package queue

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/go-redis/redis/v8"

	"github.com/khodemobin/pilo/auth/pkg/logger"
	"github.com/vmihailenco/taskq/v3"
	"github.com/vmihailenco/taskq/v3/redisq"
)

type Queue struct {
	queue  taskq.Queue
	logger logger.Logger
}

func New(rc *redis.Client, logger logger.Logger) *Queue {
	QueueFactory := redisq.NewFactory()
	q := QueueFactory.RegisterQueue(&taskq.QueueOptions{
		Name:  "api-worker",
		Redis: rc,
	})

	return &Queue{
		queue:  q,
		logger: logger,
	}
}

func (q Queue) Add(t *taskq.TaskOptions, m *taskq.Message) {
	taskq.RegisterTask(t)
	err := q.queue.Add(m)
	if err != nil {
		q.logger.Fatal(err)
	}
}

func (q Queue) Start() os.Signal {
	ch := make(chan os.Signal, 2)
	signal.Notify(
		ch,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)
	for {
		sig := <-ch
		switch sig {
		case syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM:
			return sig
		}
	}
}
