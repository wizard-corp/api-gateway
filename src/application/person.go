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
		return nil, err
	}

	return &person, nil
}

func (lu *PersonUsecase) NewPerson(person *domain.Person) error {
	return lu.PersonRepo.Create(person)
}
