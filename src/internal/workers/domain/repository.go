package domain

type WorkerRepository interface {
	CreateWorker(Worker) error
    IsWorkerAssigned(idUsuario int) (bool, error)
}
