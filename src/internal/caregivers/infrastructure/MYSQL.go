package infrastructure

import (
	"database/sql"
	"fmt"
	"pulse_sense/src/internal/caregivers/domain"
)

type CaregiverRepo struct {
	db *sql.DB
}

func NewCaregiverRepo(db *sql.DB) *CaregiverRepo {
	return &CaregiverRepo{
		db: db,
	}
}

func (r *CaregiverRepo) CreateCaregiver(Caregiver domain.Caregiver) error {
	query := `INSERT INTO cuidadores (
		idusuario, 
		idpaciente,
		turno 
	) VALUES (?, ?, ?)`

	_, err := r.db.Exec(
		query,
		Caregiver.IdUsuario,
		Caregiver.IdPaciente, 
		Caregiver.Turno,
	)
	if err != nil {
		return fmt.Errorf("error al crear Caregiver: %w", err)
	}

	return nil
}

func (r *CaregiverRepo) IsCaregiverAssigned(idUsuario int) (bool, error) { 
    var count int
    query := `SELECT COUNT(*) FROM cuidadores WHERE idusuario = ?`
    
    err := r.db.QueryRow(query, idUsuario).Scan(&count)
    if err != nil {
        return false, fmt.Errorf("error al verificar asignación: %w", err)
    }
    
    return count > 0, nil
}

func (r *CaregiverRepo) DeleteCaregiver(idUsuario int, idPaciente int) error {
	query := "DELETE FROM cuidadores WHERE idusuario = ? AND idpaciente = ?"
	result, err := r.db.Exec(query, idUsuario, idPaciente)
	if err != nil {
		return fmt.Errorf("error al eliminar cuiadador: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar eliminación: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("cuidador no encontrado")
	}

	return nil
}