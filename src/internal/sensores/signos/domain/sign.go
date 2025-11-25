package domain

type Sign struct {
	IDSignosPaciente int `json:"id_signos_paciente"`
	IDPaciente 		 int `json:"id_paciente"`
	IDSigno 		 int `json:"id_signo"`
	Valor 			 float64 `json:"valor"`
	Unidad 			 string `json:"unidad"`
	Fecha 			 string `json:"fecha"`
	Hora 			 string `json:"hora"`
	Turno 			 string `json:"turno,omitempty"` 
}