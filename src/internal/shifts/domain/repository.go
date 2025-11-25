package domain

type ShiftRepository interface {
	CreateShift(Shift) error
}
