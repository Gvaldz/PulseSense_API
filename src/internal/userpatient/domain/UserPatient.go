package domain

type UserPatientResponse struct {
	UsuarioID         int32  `json:"doctor_id"`
	UsuarioNombre     string `json:"doctor_nombre"`
	UsuarioApellidoP  string `json:"doctor_apellido_p"`
	UsuarioApellidoM  string `json:"doctor_apellido_m"`
	UsuarioCorreo     string `json:"doctor_correo"`
	PacienteID        int32  `json:"paciente_id"`
	PacienteNombre    string `json:"paciente_nombre"`
	PacienteApellidoP string `json:"paciente_apellido_p"`
	PacienteApellidoM string `json:"paciente_apellido_m"`
	Turno 			  string `json:"turno"`
}