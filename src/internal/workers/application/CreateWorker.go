package application

import (
	"pulse_sense/src/internal/workers/domain")

type CreateWorker struct {
	repo domain.WorkerRepository
}

func NewCreateWorker(repo domain.WorkerRepository) *CreateWorker {
	return &CreateWorker{repo: repo}
}

func (c *CreateWorker) Execute(Worker domain.Worker) error {
	return c.repo.CreateWorker(Worker)
}
