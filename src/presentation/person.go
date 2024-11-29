package presentation

import (
	"time"

	"github.com/wizard-corp/api-gateway/src/application"
	"github.com/wizard-corp/api-gateway/src/bootstrap"
	"github.com/wizard-corp/api-gateway/src/domain"
	"github.com/wizard-corp/api-gateway/src/mymongo"
)

type PersonController struct {
	uc *application.PersonUsecase
}

func NewPersonController(timeout time.Duration, app *bootstrap.App) *PersonController {
	repo := mymongo.NewPersonRepository(app.Mongo, timeout)
	uc := application.PersonUsecase{PersonRepo: repo}
	return &PersonController{&uc}
}

func (lc *PersonController) GetPersonByID(identifier string) (*domain.Person, error) {
	personResponse, err := lc.uc.GetPersonByID(identifier)
	if err != nil {
		return nil, err
	}

	return personResponse, nil
}

func (lc *PersonController) GetPersonByName(name string) (*domain.Person, error) {
	personResponse, err := lc.uc.GetPersonByName(name)
	if err != nil {
		return nil, err
	}

	return personResponse, nil
}

func (lc *PersonController) CreatePerson(
	firtsName string,
	patriLineal string,
	matriLineal string,
	address string,
	birthDate string) error {
	person, err := domain.NewPerson(
		firtsName,
		patriLineal,
		matriLineal,
		address,
		birthDate)
	if err != nil {
		return err
	}
	return lc.uc.CreatePerson(person)
}
