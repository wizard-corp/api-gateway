package domain

import (
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const BIRTHDATE_FORMAT = "2006-01-02"

type Person struct {
	ID          primitive.ObjectID `bson:"_id"`
	FirtsName   string             `bson:"firtsName"`
	PatriLineal string             `bson:"patriLineal"`
	MatriLineal string             `bson:"matriLineal"`
	Address     string             `bson:"address"`
	BirthDate   string             `bson:"brithDate"`
}

type PersonRepository interface {
	Create(person *Person) error
	GetPersonByID(identifier string) (Person, error)
	GetPersonByName(name string) (Person, error)
}

func (p *Person) IsValidPerson() []string {
	var errors []string

	if IsValidBirthDate(p.BirthDate) {
		errors = append(errors, INVALID_FORMAT)
	}

	return errors
}

func IsValidBirthDate(birthdateString string) bool {
	_, err := time.Parse(BIRTHDATE_FORMAT, birthdateString)
	return err != nil
}

func NewPerson(
	firtsName string,
	patriLineal string,
	matriLineal string,
	address string,
	birthDate string) (*Person, error) {
	person := Person{
		ID:          primitive.NewObjectID(),
		FirtsName:   firtsName,
		PatriLineal: patriLineal,
		MatriLineal: matriLineal,
		Address:     address,
		BirthDate:   birthDate}
	errs := person.IsValidPerson()
	if len(errs) > 0 {
		return nil, NewDomainError(INVALID_SCHEMA, strings.Join(errs, "\n"))
	}
	return &person, nil
}
