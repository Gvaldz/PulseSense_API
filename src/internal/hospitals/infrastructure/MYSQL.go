package infrastructure

import (
	"database/sql"
	"fmt"
	"pulse_sense/src/internal/hospitals/domain"
)

type HospitalRepo struct {
	db *sql.DB
}

func (r *HospitalRepo) GetHospitalByUser(iduser int32) ([]domain.Hospital, error) {
	panic("unimplemented")
}

func NewHospitalRepo(db *sql.DB) *HospitalRepo {
	return &HospitalRepo{
		db: db,
	}
}

func (r *HospitalRepo) CreateHospital(hospital domain.Hospital) error {
	query := `INSERT INTO hospitales (
		nombre, 
		ubicacion, 
		clues
	) VALUES (?, ?, ?)`

	_, err := r.db.Exec(
		query,
		hospital.Nombre,
		hospital.Ubicacion,
		hospital.Clues,
	)
	if err != nil {
		return fmt.Errorf("error al crear hospital: %w", err)
	}

	return nil
}

func (r *HospitalRepo) GetAllHospital() ([]domain.Hospital, error) {
	query := `SELECT 
		idhospital, 
		nombre, 
		ubicacion, 
		clues
	FROM hospitales`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener hospitales: %w", err)
	}
	defer rows.Close()

	var hospitals []domain.Hospital
	for rows.Next() {
		var hospital domain.Hospital
		if err := rows.Scan(
			&hospital.IdHospital,
			&hospital.Nombre,
			&hospital.Ubicacion,
			&hospital.Clues,
		); err != nil {
			return nil, fmt.Errorf("error al escanear hospital: %w", err)
		}
		hospitals = append(hospitals, hospital)
	}

	return hospitals, nil
}

func (r *HospitalRepo) GetHospitalByID(idHospital string) (domain.Hospital, error) {
	var hospital domain.Hospital
	query := `SELECT 
		idhospital, 
		nombre, 
		ubicacion, 
		clues, 
	FROM hospitales WHERE idhospital = ?`

	err := r.db.QueryRow(query, idHospital).Scan(
		&hospital.IdHospital,
		&hospital.Nombre,
		&hospital.Ubicacion,
		&hospital.Clues,
	)
	if err != nil {
		return hospital, fmt.Errorf("error al obtener hospital: %w", err)
	}
	return hospital, nil
}

func (r *HospitalRepo) UpdateHospital(id string, hospital domain.Hospital) error {
	query := `UPDATE hospitales SET 
		nombre = ?, 
		ubicacion = ?, 
		clues = ?, 
	WHERE idhospitales = ?`

	result, err := r.db.Exec(
		query,
		hospital.Nombre,
		hospital.Ubicacion,
		hospital.Clues,
		id,
	)
	if err != nil {
		return fmt.Errorf("error al actualizar hospital: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar actualización: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("hospital no encontrado")
	}

	return nil
}

func (r *HospitalRepo) SearchHospital(criteria string) ([]domain.Hospital, error) {
	query := `SELECT 
		idhospital, 
		nombre, 
		ubicacion, 
		clues
	FROM hospitales 
	WHERE nombre LIKE ? OR ubicacion LIKE ? OR clues LIKE ?`

	searchTerm := "%" + criteria + "%"
	rows, err := r.db.Query(query, searchTerm, searchTerm, searchTerm)
	if err != nil {
		return nil, fmt.Errorf("error al buscar hospitales: %w", err)
	}
	defer rows.Close()

	var hospitales []domain.Hospital
	for rows.Next() {
		var hospital domain.Hospital
		if err := rows.Scan(
			&hospital.IdHospital,
			&hospital.Nombre,
			&hospital.Ubicacion,
			&hospital.Clues,
		); err != nil {
			return nil, fmt.Errorf("error al escanear hospital: %w", err)
		}
		hospitales = append(hospitales, hospital)
	}

	return hospitales, nil
}
