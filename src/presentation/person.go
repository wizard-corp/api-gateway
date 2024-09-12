package presentation

import (
	"errors"
	"strings"
	"time"

	"github.com/wizard-corp/api-gateway/src/application"
	"github.com/wizard-corp/api-gateway/src/bootstrap"
	"github.com/wizard-corp/api-gateway/src/domain"
	"github.com/wizard-corp/api-gateway/src/mymongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (lc *PersonController) NewPerson(
	firtsName string,
	patriLineal string,
	matriLineal string,
	address string,
	birthDate string) error {
	person := domain.Person{
		ID:          primitive.NewObjectID(),
		FirtsName:   firtsName,
		PatriLineal: patriLineal,
		MatriLineal: matriLineal,
		Address:     address,
		BirthDate:   birthDate}
	errs := person.IsValidPerson()
	if len(errs) > 0 {
		txt := domain.INVALID_SCHEMA + "\n" + strings.Join(errs, "\n")
		return errors.New(txt)
	}

	return lc.uc.NewPerson(&person)
}
