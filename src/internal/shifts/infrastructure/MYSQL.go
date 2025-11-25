package infrastructure

import (
	"database/sql"
	"fmt"
	"pulse_sense/src/internal/shifts/domain"
)

type ShiftRepo struct {
	db *sql.DB
}

func NewShiftRepo(db *sql.DB) *ShiftRepo {
	return &ShiftRepo{
		db: db,
	}
}

func (r *ShiftRepo) CreateShift(Shift domain.Shift) error {
	query := `INSERT INTO turnos (
		idusuario, 
		turno 
	) VALUES (?, ?)`

	_, err := r.db.Exec(
		query,
		Shift.IdUsuario,
		Shift.Turno,
	)
	if err != nil {
		return fmt.Errorf("error al crear turno: %w", err)
	}

	return nil
}
