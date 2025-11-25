package domain

type MotionRepository interface {
	GetByPatient(IDPatient int) ([]Motion, error)
	CreateMotion(Motion) error
}
