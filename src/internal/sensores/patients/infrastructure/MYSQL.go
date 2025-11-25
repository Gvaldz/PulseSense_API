package infrastructure

import (
	"database/sql"
	"fmt"
	"pulse_sense/src/internal/sensores/patients/domain"
)

type PatientRepo struct {
	db *sql.DB
}

func (r *PatientRepo) GetAllPatient() ([]domain.Patient, error) {
	panic("unimplemented")
}

func NewPatientRepo(db *sql.DB) *PatientRepo {
	return &PatientRepo{
		db: db,
	}
}

func (r *PatientRepo) GetPatientByNurse(iduser int32) ([]domain.Patient, error) {
	query := `SELECT 
		p.idpaciente, 
		p.nombres, 
		p.apellido_p, 
		p.apellido_m, 
		p.nacimiento, 
		p.peso, 
		p.estatura, 
		p.sexo, 
		p.idtipo_de_sangre, 
		p.numero_emergencia
	FROM 
		pacientes p
	INNER JOIN 
		cuidadores c ON p.idpaciente = c.idpaciente
	INNER JOIN
		usuarios u ON c.idusuario = u.idusuario
	WHERE 
		c.idusuario = ? 
		AND u.idtipo_usuario = 2`

	rows, err := r.db.Query(query, iduser)
	if err != nil {
		return nil, fmt.Errorf("error al obtener pacientes por usuario: %w", err)
	}
	defer rows.Close()

	var patients []domain.Patient
	for rows.Next() {
		var patient domain.Patient
		if err := rows.Scan(
			&patient.IdPaciente,
			&patient.Nombres,
			&patient.Apellido_p,
			&patient.Apellido_m,
			&patient.Nacimiento,
			&patient.Peso,
			&patient.Estatura,
			&patient.Sexo,
			&patient.IDTipoSangre,
			&patient.Numero_emergencia,
		); err != nil {
			return nil, fmt.Errorf("error al escanear paciente: %w", err)
		}
		patients = append(patients, patient)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error después de iterar pacientes: %w", err)
	}

	return patients, nil
}


func (r *PatientRepo) GetPatientByUser(iduser int32) ([]domain.Patient, error) {
	query := `SELECT 
		idpaciente, 
		nombres, 
		apellido_p, 
		apellido_m, 
		nacimiento, 
		peso, 
		estatura, 
		sexo, 
		idtipo_de_sangre, 
		numero_emergencia 
	FROM pacientes
	WHERE id_doctor = ?`

	rows, err := r.db.Query(query, iduser)
	if err != nil {
		return nil, fmt.Errorf("error al obtener pacientes por usuario: %w", err)
	}
	defer rows.Close()

	var patients []domain.Patient
	for rows.Next() {
		var patient domain.Patient
		if err := rows.Scan(
			&patient.IdPaciente,
			&patient.Nombres,
			&patient.Apellido_p,
			&patient.Apellido_m,
			&patient.Nacimiento,
			&patient.Peso,
			&patient.Estatura,
			&patient.Sexo,
			&patient.IDTipoSangre,
			&patient.Numero_emergencia,
		); err != nil {
			return nil, fmt.Errorf("error al escanear paciente: %w", err)
		}
		patients = append(patients, patient)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error después de iterar pacientes: %w", err)
	}

	return patients, nil
}

func (r *PatientRepo) CreatePatient(patient domain.Patient) (int64, error) {
    query := `INSERT INTO pacientes (
        nombres, 
        apellido_p, 
        apellido_m, 
        nacimiento, 
        peso, 
        estatura, 
        sexo, 
        idtipo_de_sangre, 
        numero_emergencia,
        id_doctor
    ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

    result, err := r.db.Exec(
        query,
        patient.Nombres,
        patient.Apellido_p,
        patient.Apellido_m,
        patient.Nacimiento,
        patient.Peso,
        patient.Estatura,
        patient.Sexo,
        patient.IDTipoSangre,
        patient.Numero_emergencia,
        patient.IDDoctor,
    )
    if err != nil {
        return 0, fmt.Errorf("error al crear paciente: %w", err)
    }

    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("error al obtener último ID insertado: %w", err)
    }

    return id, nil
}

func (r *PatientRepo) GetAllPatients() ([]domain.Patient, error) {
	query := `SELECT 
		idpaciente, 
		nombres, 
		apellido_p, 
		apellido_m, 
		nacimiento, 
		peso, 
		estatura, 
		sexo, 
		idtipo_de_sangre, 
		numero_emergencia,
		id_doctor 
	FROM pacientes`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener pacientes: %w", err)
	}
	defer rows.Close()

	var patients []domain.Patient
	for rows.Next() {
		var patient domain.Patient
		if err := rows.Scan(
			&patient.IdPaciente,
			&patient.Nombres,
			&patient.Apellido_p,
			&patient.Apellido_m,
			&patient.Nacimiento,
			&patient.Peso,
			&patient.Estatura,
			&patient.Sexo,
			&patient.IDTipoSangre,
			&patient.Numero_emergencia,
			&patient.IDDoctor,
		); err != nil {
			return nil, fmt.Errorf("error al escanear paciente: %w", err)
		}
		patients = append(patients, patient)
	}

	return patients, nil
}

func (r *PatientRepo) GetPatientByID(idPaciente string) (domain.Patient, error) {
	var patient domain.Patient
	query := `SELECT 
		idpaciente, 
		nombres, 
		apellido_p, 
		apellido_m, 
		nacimiento, 
		peso, 
		estatura, 
		sexo, 
		idtipo_de_sangre, 
		numero_emergencia,
		id_doctor
	FROM pacientes WHERE idpaciente = ?`

	err := r.db.QueryRow(query, idPaciente).Scan(
		&patient.IdPaciente,
		&patient.Nombres,
		&patient.Apellido_p,
		&patient.Apellido_m,
		&patient.Nacimiento,
		&patient.Peso,
		&patient.Estatura,
		&patient.Sexo,
		&patient.IDTipoSangre,
		&patient.Numero_emergencia,
		&patient.IDDoctor,
	)
	if err != nil {
		return patient, fmt.Errorf("error al obtener paciente: %w", err)
	}
	return patient, nil
}

func (r *PatientRepo) UpdatePatient(id string, patient domain.Patient) error {
	query := `UPDATE pacientes SET 
		nombres = ?, 
		apellido_p = ?, 
		apellido_m = ?, 
		nacimiento = ?, 
		peso = ?, 
		estatura = ?, 
		sexo = ?, 
		idtipo_de_sangre = ?, 
		numero_emergencia = ? 
	WHERE idpaciente = ?`

	result, err := r.db.Exec(
		query,
		patient.Nombres,
		patient.Apellido_p,
		patient.Apellido_m,
		patient.Nacimiento,
		patient.Peso,
		patient.Estatura,
		patient.Sexo,
		patient.IDTipoSangre,
		patient.Numero_emergencia,
		id,
	)
	if err != nil {
		return fmt.Errorf("error al actualizar paciente: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar actualización: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("paciente no encontrado")
	}

	return nil
}

func (r *PatientRepo) SearchPatients(criteria string) ([]domain.Patient, error) {
	query := `SELECT 
		idpaciente, 
		nombres, 
		apellido_p, 
		apellido_m, 
		nacimiento, 
		peso, 
		estatura, 
		sexo, 
		idtipo_de_sangre, 
		numero_emergencia 
	FROM pacientes 
	WHERE nombres LIKE ? OR apellido_p LIKE ? OR apellido_m LIKE ?`

	searchTerm := "%" + criteria + "%"
	rows, err := r.db.Query(query, searchTerm, searchTerm, searchTerm)
	if err != nil {
		return nil, fmt.Errorf("error al buscar pacientes: %w", err)
	}
	defer rows.Close()

	var patients []domain.Patient
	for rows.Next() {
		var patient domain.Patient
		if err := rows.Scan(
			&patient.IdPaciente,
			&patient.Nombres,
			&patient.Apellido_p,
			&patient.Apellido_m,
			&patient.Nacimiento,
			&patient.Peso,
			&patient.Estatura,
			&patient.Sexo,
			&patient.IDTipoSangre,
			&patient.Numero_emergencia,
		); err != nil {
			return nil, fmt.Errorf("error al escanear paciente: %w", err)
		}
		patients = append(patients, patient)
	}

	return patients, nil
}
