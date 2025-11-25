package domain

type CaregiverRepository interface {
	CreateCaregiver(Caregiver) error
    IsCaregiverAssigned(idUsuario int) (bool, error)
	DeleteCaregiver(idUsuario int, IdPaciente int) error
}
