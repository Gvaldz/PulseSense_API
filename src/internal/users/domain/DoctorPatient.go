package domain

type UserPatientResponse struct {
	DoctorID       int32  `json:"doctor_id"`
	DoctorNombre   string `json:"doctor_nombre"`
	DoctorApellido string `json:"doctor_apellido"`
	DoctorCorreo   string `json:"doctor_correo"`
	PacienteID     int32  `json:"paciente_id"`
	PacienteNombre string `json:"paciente_nombre"`
}
