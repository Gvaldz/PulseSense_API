package domain

type PatientRepository interface {
	CreatePatient(Patient Patient) (int64, error)
	GetAllPatient() ([]Patient, error)
	GetPatientByID(idPatient string) (Patient, error)
	GetPatientByUser(iduser int32) ([]Patient, error)
	GetPatientByNurse(iduser int32) ([]Patient, error)
	UpdatePatient(idPatient string, Patient Patient) error
}
