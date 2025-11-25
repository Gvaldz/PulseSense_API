package infrastructure

import (
	"database/sql"
	"fmt"
	amqpConsumer "pulse_sense/src/consumer_amqp"
	"pulse_sense/src/internal/sensores/signos/domain"
)

type SignsRepo struct {
	db           *sql.DB
	amqpConsumer *amqpConsumer.RabbitMQConsumer
}

func (r *SignsRepo) GetByPatient(IDPaciente int) ([]domain.Sign, error) {
    return r.GetSignsByPatient(IDPaciente)
}

func NewSignsRepo(db *sql.DB, amqpConsumer *amqpConsumer.RabbitMQConsumer) *SignsRepo {
	return &SignsRepo{
		db:           db,
		amqpConsumer: amqpConsumer,
	}
}

func (r *SignsRepo) CreateSigns(sign domain.Sign) error {
	query := `INSERT INTO signos_paciente (
		idpaciente, 
		idsigno, 
		valor, 
		unidad
	) VALUES (?, ?, ?, ?)`

	_, err := r.db.Exec(
		query,
		sign.IDPaciente,
		sign.IDSigno,
		sign.Valor,
		sign.Unidad,
	)
	if err != nil {
		return fmt.Errorf("error al guardar signo vital: %w", err)
	}

	return nil
}

func (r *SignsRepo) GetSignsByPatient(IDPaciente int) ([]domain.Sign, error) {
	query := `SELECT 
		idsignos_paciente,
		idpaciente,
		idsigno,
		valor,
		unidad,
		fecha,
		hora
	FROM signos_paciente WHERE idpaciente = ?`

	rows, err := r.db.Query(query, IDPaciente)
	if err != nil {
		return nil, fmt.Errorf("error al obtener signos vitales: %w", err)
	}
	defer rows.Close()

	var signs []domain.Sign
	for rows.Next() {
		var sign domain.Sign
		if err := rows.Scan(
			&sign.IDSignosPaciente,
			&sign.IDPaciente,
			&sign.IDSigno,
			&sign.Valor,
			&sign.Unidad,
			&sign.Fecha,
			&sign.Hora,
		); err != nil {
			return nil, fmt.Errorf("error al escanear signo vital: %w", err)
		}
		signs = append(signs, sign)
	}

	return signs, nil
}

func (r *SignsRepo) GetSignsByTypeAndTimeRange(IDPaciente int, IDSigno int, fecha string, turno string) ([]domain.Sign, error) {
    query := `SELECT 
        idsignos_paciente,
        idpaciente,
        idsigno,
        valor,
        unidad,
        fecha,
        hora,
        turno
    FROM signos_paciente 
    WHERE idpaciente = ? 
    AND idsigno = ?
    AND fecha = ?
    AND LOWER(turno) = LOWER(?)`

    fmt.Println("Query a ejecutar:", query)
    
    rows, err := r.db.Query(query, IDPaciente, IDSigno, fecha, turno)
    if err != nil {
        fmt.Println("Error al ejecutar query:", err)
        return nil, fmt.Errorf("error al obtener signos por tipo y turno: %w", err)
    }
    defer rows.Close()

    var signs []domain.Sign
    for rows.Next() {
        var sign domain.Sign
        if err := rows.Scan(
            &sign.IDSignosPaciente,
            &sign.IDPaciente,
            &sign.IDSigno,
            &sign.Valor,
            &sign.Unidad,
            &sign.Fecha,
            &sign.Hora,
            &sign.Turno,
        ); err != nil {
            fmt.Println("Error al escanear fila:", err)
            return nil, fmt.Errorf("error al escanear signo vital: %w", err)
        }
        fmt.Printf("Registro encontrado: %+v\n", sign)
        signs = append(signs, sign)
    }

    if len(signs) == 0 {
        fmt.Println("No se encontraron registros con los criterios especificados")
    }

    return signs, nil
}

func (r *SignsRepo) GetLatestSigns(IDPaciente int, IDSigno float64, limit int) ([]domain.Sign, error) {
	query := `SELECT 
		idsignos_paciente,
		idpaciente,
		idsigno,
		valor,
		unidad,
		fecha,
		hora
	FROM signos_paciente 
	WHERE idpaciente = ? AND idsigno = ?
	ORDER BY fecha DESC, hora DESC
	LIMIT ?`

	rows, err := r.db.Query(query, IDPaciente, IDSigno, limit)
	if err != nil {
		return nil, fmt.Errorf("error al obtener últimos signos vitales: %w", err)
	}
	defer rows.Close()

	var signs []domain.Sign
	for rows.Next() {
		var sign domain.Sign
		if err := rows.Scan(
			&sign.IDSignosPaciente,
			&sign.IDPaciente,
			&sign.IDSigno,
			&sign.Valor,
			&sign.Unidad,
			&sign.Fecha,
			&sign.Hora,
		); err != nil {
			return nil, fmt.Errorf("error al escanear signo vital: %w", err)
		}
		signs = append(signs, sign)
	}

	return signs, nil
}
