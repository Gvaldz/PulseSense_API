package consumeramqp

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"pulse_sense/src/core"

	dependencesMotion "pulse_sense/src/internal/sensores/motion/domain"
	controllersMotion "pulse_sense/src/internal/sensores/motion/infrastructure/controllers"
	dependencesPatient "pulse_sense/src/internal/sensores/patients/domain"
	controllersPatient "pulse_sense/src/internal/sensores/patients/infrastructure/controllers"
	dependencesSigns "pulse_sense/src/internal/sensores/signos/domain"
	controllersSigns "pulse_sense/src/internal/sensores/signos/infrastructure/controllers"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQConsumer struct {
	conn          *core.AMQPConnection
	CreateSign    *controllersSigns.CreateSignsController
	CreatePatient *controllersPatient.CreatePatientController
	CreateMotion  *controllersMotion.CreateMotionController
}

func NewRabbitMQConsumer(
	conn *core.AMQPConnection,
	createSign *controllersSigns.CreateSignsController,
	createPatient *controllersPatient.CreatePatientController,
	createMotion *controllersMotion.CreateMotionController,
) *RabbitMQConsumer {
	return &RabbitMQConsumer{
		conn:          conn,
		CreateSign:    createSign,
		CreatePatient: createPatient,
		CreateMotion:  createMotion,
	}
}

func (c *RabbitMQConsumer) Start() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error cargando .env: %v", err)
	}

	amqpServer := os.Getenv("AMQP_SERVER")

	connRabbit, err := amqp.Dial(amqpServer)
	if err != nil {
		log.Fatalf("Error conectando a RabbitMQ: %v", err)
	}
	log.Println("Conexión a RabbitMQ establecida.")
	defer connRabbit.Close()

	ch, err := connRabbit.Channel()
	if err != nil {
		log.Fatalf("Error abriendo canal en RabbitMQ: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("sensores", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Error declarando cola: %v", err)
	}

	err = ch.QueueBind(
		"sensores",        // queue name
		"sensores/signos", // routing key
		"amq.topic",       // exchange
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Error al enlazar cola a topic MQTT: %v", err)
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Error al consumir mensajes: %v", err)
	}
	log.Println("Esperando mensajes...")

	for msg := range msgs {
		var sensorData struct {
			Sensor     string  `json:"sensor"`
			IDPaciente int     `json:"IDPaciente"`
			Signo      int     `json:"Signo,omitempty"`
			Valor      float64 `json:"Valor,omitempty"`
			Unidad     string  `json:"Unidad,omitempty"`
			Movimiento int     `json:"movimiento,omitempty"`
		}

		if err := json.Unmarshal(msg.Body, &sensorData); err != nil {
			log.Printf("Error deserializando el mensaje: %v", err)
			continue
		}

		log.Printf("Mensaje recibido: Sensor: %s, IDPaciente: %v\n", sensorData.Sensor, sensorData.IDPaciente)

		if sensorData.Sensor == "idpatient" {

			if sensorData.IDPaciente == 0 {
				log.Println("Advertencia: ID no valido.")
				continue
			}

			paciente := dependencesPatient.Patient{
				IdPaciente: int32(sensorData.IDPaciente),
			}

			if err := c.CreatePatient.Processpatient(paciente); err != nil {
				log.Printf("Error al procesar paciente: %v", err)
			}
		}

		if sensorData.Sensor == "IDPaciente" {
			if sensorData.IDPaciente == 0 {
				log.Println("Advertencia: ID no valido.")
				continue
			}

			var idPacienteInt int32 = int32(sensorData.IDPaciente)

			paciente := dependencesPatient.Patient{
				IdPaciente: idPacienteInt,
			}

			if err := c.CreatePatient.Processpatient(paciente); err != nil {
				log.Printf("Error al procesar paciente: %v", err)
			}
		}

		switch sensorData.Sensor {
		case "signos":
			if sensorData.Signo == 0 {
				log.Printf("Signo no valido para el paciente ID: %s\n", sensorData.IDPaciente)
				continue
			}
			log.Printf("Procesando Signo: %v", sensorData.Signo)

			idPacienteInt := sensorData.IDPaciente
			sign := dependencesSigns.Sign{
				IDPaciente: idPacienteInt,
				IDSigno:    sensorData.Signo,
				Valor:      sensorData.Valor,
				Unidad:     sensorData.Unidad,
			}

			log.Printf("Signo procesado: %v", sign)

			if c.CreateSign != nil {
				err := c.CreateSign.ProcessSign(sign)
				if err != nil {
					log.Printf("Error procesando Signo: %v", err)
				} else {
					log.Println("Signo procesado exitosamente.")
				}
			} else {
				log.Println("El controlador de Signo es nil, no se puede procesar.")
			}

		case "movimiento":
			if sensorData.IDPaciente == 0 {
				log.Println("Advertencia: ID de paciente no válido para movimiento.")
				continue
			}

			mov := sensorData.Movimiento == 1
			fmt.Printf("Mensaje de Movimiento recibido: %+v\n", sensorData)

			motion := dependencesMotion.Motion{
				IDPaciente: sensorData.IDPaciente,
				Movimiento: mov,
			}

			if c.CreateMotion != nil {
				if err := c.CreateMotion.ProcessMotion(motion); err != nil {
					log.Printf("Error al procesar el movimiento: %v", err)
				} else {
					log.Println("Movimiento procesado exitosamente.")
				}
			} else {
				log.Println("El controlador de Movimiento es nil, no se puede procesar.")
			}
		}
	}
}
