package domain

import (
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
	GetPersonByID(id string) (Person, error)
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
