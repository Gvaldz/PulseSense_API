package infrastructure

import (
	"database/sql"
	"fmt"
	amqpConsumer "pulse_sense/src/consumer_amqp"
	"pulse_sense/src/internal/sensores/motion/domain"
)

type MotionRepo struct {
	db           *sql.DB
	amqpConsumer *amqpConsumer.RabbitMQConsumer
}

func NewMotionRepo(db *sql.DB, amqpConsumer *amqpConsumer.RabbitMQConsumer) *MotionRepo {
	return &MotionRepo{
		db:           db,
		amqpConsumer: amqpConsumer,
	}
}

func (r *MotionRepo) CreateMotion(motion domain.Motion) error {
	query := "INSERT INTO movimientos (idpaciente, movimiento, hora) VALUES (?, ?, NOW())"
	_, err := r.db.Exec(query, motion.IDPaciente, motion.Movimiento)
	if err != nil {
		return fmt.Errorf("error al guardar movimiento: %w", err)
	}

	return nil
}

func (r *MotionRepo) GetByPatient(IDPatient int) ([]domain.Motion, error) {
	query := "SELECT idmovimiento, idpaciente, movimiento, hora FROM movimientos WHERE idpaciente = ?"
	rows, err := r.db.Query(query, IDPatient)
	if err != nil {
		return nil, fmt.Errorf("error al obtener movimientos: %w", err)
	}
	defer rows.Close()

	var motions []domain.Motion
	for rows.Next() {
		var motion domain.Motion
		if err := rows.Scan(&motion.IDMovimiento, &motion.IDPaciente, &motion.Movimiento, &motion.HoraRegistro); err != nil {
			return nil, fmt.Errorf("error al escanear movimiento: %w", err)
		}
		motions = append(motions, motion)
	}

	return motions, nil
}
