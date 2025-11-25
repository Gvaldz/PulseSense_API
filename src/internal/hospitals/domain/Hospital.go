package domain

type Hospital struct {
	IdHospital        int32  `json:"id_hospital"`
	Nombre            string `json:"nombre"`
	Ubicacion         string `json:"ubicacion"`
	Clues	          string `json:"clues"`
}
