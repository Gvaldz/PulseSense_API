package infrastructure

import (
	"database/sql"
	"fmt"
	"pulse_sense/src/internal/workers/domain"
)

type WorkerRepo struct {
	db *sql.DB
}

func NewWorkerRepo(db *sql.DB) *WorkerRepo {
	return &WorkerRepo{
		db: db,
	}
}

func (r *WorkerRepo) CreateWorker(Worker domain.Worker) error {
	query := `INSERT INTO trabajadores (
		idusuario, 
		idhospital 
	) VALUES (?, ?)`

	_, err := r.db.Exec(
		query,
		Worker.IdUsuario,
		Worker.IdHospital,
	)
	if err != nil {
		return fmt.Errorf("error al crear Worker: %w", err)
	}

	return nil
}

func (r *WorkerRepo) IsWorkerAssigned(idUsuario int) (bool, error) { 
    var count int
    query := `SELECT COUNT(*) FROM trabajadores WHERE idusuario = ?`
    
    err := r.db.QueryRow(query, idUsuario).Scan(&count)
    if err != nil {
        return false, fmt.Errorf("error al verificar asignación: %w", err)
    }
    
    return count > 0, nil
}