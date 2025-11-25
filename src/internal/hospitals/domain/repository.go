package domain

type HospitalRepository interface {
	CreateHospital(Hospital) error
	GetAllHospital() ([]Hospital, error)
	GetHospitalByID(idHospital string) (Hospital, error)
	GetHospitalByUser(iduser int32) ([]Hospital, error)
	UpdateHospital(idHospital string, Hospital Hospital) error
	SearchHospital(name string) ([]Hospital, error)
}
