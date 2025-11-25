package domain

type User struct {
	IdUsuario  int32  `json:"iduser"`
	Nombres    string `json:"nombres"`
	Apellido_m  string `json:"apellidos"`
	Apellido_p string `json:"apellido_p"`
	Correo     string `json:"correo"`
	Contrasena string `json:"contrasena"`
	Tipo       int32 `json:"tipo"`
	FCMToken   string `json:"fcm_token"`
}
