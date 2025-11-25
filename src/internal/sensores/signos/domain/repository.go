package domain

type SignRepository interface {
    GetByPatient(IDPaciente int) ([]Sign, error)
	CreateSigns(Sign) error
	GetSignsByTypeAndTimeRange(IDPaciente int, IDtipo int, fecha string, turno string) ([]Sign, error)
}
