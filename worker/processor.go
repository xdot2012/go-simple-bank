package worker

import (
	"context"

	"github.com/hibiken/asynq"
	db "github.com/xdot2012/simple-bank/db/sqlc"
)

type TaskProcessor interface {
	Start() error
	ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error
}

type RedisTaskProcessor struct {
	server *asynq.Server
	store  db.Store
}

const (
	QueueCritical = "critical"
	QueueDefault  = "default"
	QueueLow      = "low"
)

func NewRedisTaskProcessor(redisOpt asynq.RedisClientOpt, store db.Store) TaskProcessor {
	server := asynq.NewServer(
		redisOpt,
		asynq.Config{
			Queues: map[string]int{
				QueueCritical: 10,
				QueueDefault:  5,
				QueueLow:      1,
			},
		},
	)

	return &RedisTaskProcessor{
		server: server,
		store:  store,
	}
}

func (processor *RedisTaskProcessor) Start() error {
	mux := asynq.NewServeMux()

	// Register tasks here
	mux.HandleFunc(TaskSendVerifyEmail, processor.ProcessTaskSendVerifyEmail)

	return processor.server.Start(mux)
}
