package domain

type Patient struct {
	IdPaciente        int32   `json:"id_paciente"`
	Nombres           string  `json:"nombres"`
	Apellido_p        string  `json:"apellido_p"`
	Apellido_m        string  `json:"apellido_m"`
	Nacimiento        string  `json:"nacimiento"`
	Peso              float32 `json:"peso"`
	Estatura          float32 `json:"estatura"`
	Sexo              string  `json:"sexo"`
	IDTipoSangre      int	  `json:"id_tipo_sangre"`
	Numero_emergencia string  `json:"numero_emergencia"`
	IDDoctor          int32   `json:"id_doctor"`
}
