package domain

type DoctorPatientRepository interface {
	GetDoctorsByPatientId(idCuidador int32) ([]UserPatientResponse, error)
	GetNursesByPatientId(idDoctor int32) ([]UserPatientResponse, error)
}
