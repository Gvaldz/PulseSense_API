package application

import (
	"pulse_sense/src/internal/workers/domain")

type CheckWorkerAssignment struct {
    repo domain.WorkerRepository
}

func NewCheckWorkerAssignment(repo domain.WorkerRepository) *CheckWorkerAssignment {
    return &CheckWorkerAssignment{repo: repo}
}

func (c *CheckWorkerAssignment) Execute(idUsuario int) (bool, error) {
    return c.repo.IsWorkerAssigned(idUsuario) 
}