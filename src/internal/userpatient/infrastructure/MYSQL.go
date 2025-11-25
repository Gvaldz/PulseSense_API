package infrastructure

import (
	"database/sql"
	"fmt"
	"pulse_sense/src/internal/userpatient/domain"
)

type UserPatientRepo struct {
	db *sql.DB
}

func NewUserPatientRepo(db *sql.DB) *UserPatientRepo {
	return &UserPatientRepo{db: db}
}

func (r *UserPatientRepo) GetDoctorsByPatientId(idCuidador int32) ([]domain.UserPatientResponse, error) {
	query := `
	SELECT 
		d.idusuario,
		d.nombres,
		d.apellido_p,
		d.apellido_m,
		d.correo,
		p.idpaciente,
		p.nombres,
		p.apellido_p,
		p.apellido_m
	FROM cuidadores c
	JOIN pacientes p ON c.idpaciente = p.idpaciente
	JOIN usuarios d ON p.id_doctor = d.idusuario
	WHERE c.idusuario = ?`

	rows, err := r.db.Query(query, idCuidador)
	if err != nil {
		return nil, fmt.Errorf("error al obtener datos: %w", err)
	}
	defer rows.Close()

	var results []domain.UserPatientResponse
	for rows.Next() {
		var item domain.UserPatientResponse
		if err := rows.Scan(
			&item.UsuarioID,
			&item.UsuarioNombre,
			&item.UsuarioApellidoP,
			&item.UsuarioApellidoM,
			&item.UsuarioCorreo,
			&item.PacienteID,
			&item.PacienteNombre,
			&item.PacienteApellidoP,
			&item.PacienteApellidoM,
		); err != nil {
			return nil, fmt.Errorf("error al escanear resultado: %w", err)
		}
		results = append(results, item)
	}

	return results, nil
}

func (r *UserPatientRepo) GetNursesByPatientId(idDoctor int32) ([]domain.UserPatientResponse, error) {
    query := `
    SELECT 
        n.idusuario,
        n.nombres,
        n.apellido_p,
        n.apellido_m,
        n.correo,
        p.idpaciente,
        p.nombres,
        p.apellido_p,
        p.apellido_m,
		c.turno
    FROM pacientes p
    JOIN cuidadores c ON p.idpaciente = c.idpaciente
    JOIN usuarios n ON c.idusuario = n.idusuario
    WHERE p.id_doctor = ?`

    rows, err := r.db.Query(query, idDoctor)
    if err != nil {
        return nil, fmt.Errorf("error al obtener enfermeros por doctor: %w", err)
    }
    defer rows.Close()

    var results []domain.UserPatientResponse
    for rows.Next() {
        var item domain.UserPatientResponse
        if err := rows.Scan(
            &item.UsuarioID,
            &item.UsuarioNombre,
            &item.UsuarioApellidoP,
            &item.UsuarioApellidoM,
            &item.UsuarioCorreo,
            &item.PacienteID,
            &item.PacienteNombre,
            &item.PacienteApellidoP,
            &item.PacienteApellidoM,
			&item.Turno, 
        ); err != nil {
            return nil, fmt.Errorf("error al escanear resultado: %w", err)
        }
        results = append(results, item)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error después de iterar resultados: %w", err)
    }

    return results, nil
}