package application

import (
	"github.com/wizard-corp/api-gateway/src/domain"
)

type PersonUsecase struct {
	PersonRepo domain.PersonRepository
}

func (lu *PersonUsecase) GetPersonByID(perdonId string) (*domain.Person, error) {
	person, err := lu.PersonRepo.GetPersonByID(perdonId)
	if err != nil {
		return nil, domain.NewInfrastructureError(domain.DB_FETCH_BY_PARAM, err.Error())
	}

	return &person, nil
}

func (lu *PersonUsecase) GetPersonByName(name string) (*domain.Person, error) {
	person, err := lu.PersonRepo.GetPersonByName(name)
	if err != nil {
		return nil, domain.NewInfrastructureError(domain.DB_FETCH_BY_PARAM, err.Error())
	}

	return &person, nil
}

func (lu *PersonUsecase) CreatePerson(person *domain.Person) error {
	err := lu.PersonRepo.Create(person)
	if err != nil {
		return domain.NewInfrastructureError(domain.DB_INSERT_FAILED, err.Error())
	}
	return nil
}
