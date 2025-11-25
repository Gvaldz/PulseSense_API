package domain

type Motion struct {
	IDMovimiento int `json:"idmovimiento"`
	IDPaciente int `json:"idpaciente"`
	Movimiento bool `json:"movimiento"`
	HoraRegistro string `json:"hora_registro"`
}