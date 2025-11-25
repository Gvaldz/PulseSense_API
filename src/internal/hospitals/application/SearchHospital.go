package application

import (
	"pulse_sense/src/internal/hospitals/domain")

type SearchHospital struct {
	repo domain.HospitalRepository
}

func NewSearchHospital(repo domain.HospitalRepository) *SearchHospital {
	return &SearchHospital{repo: repo}
}

func (cp *SearchHospital) Execute(name string) ([]domain.Hospital, error) {
	return cp.repo.SearchHospital(name)
}
