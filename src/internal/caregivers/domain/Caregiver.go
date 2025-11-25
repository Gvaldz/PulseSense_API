package domain

type Caregiver struct {
	IdUsuario 		  int32  `json:"id_usuario"`
	IdPaciente        int32  `json:"id_paciente"`
	Turno			  string `json:"turno"`
}
