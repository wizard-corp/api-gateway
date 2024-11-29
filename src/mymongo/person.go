package mymongo

import (
	"context"
	"errors"
	"time"

	"github.com/wizard-corp/api-gateway/src/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoPerson struct {
	personRepo     *mongo.Collection
	ContextTimeout time.Duration
}

func NewPersonRepository(mongodb *MongoDB, timeout time.Duration) *MongoPerson {
	collection := mongodb.Mdb.Collection("person")
	return &MongoPerson{collection, timeout}
}

func (uc *MongoPerson) GetPersonByID(id string) (domain.Person, error) {
	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Person{}, err
	}
	filter := bson.D{{Key: "_id", Value: idHex}}
	return uc.GetSinglePersonByParam(filter)
}

func (uc *MongoPerson) GetPersonByName(name string) (domain.Person, error) {
	filter := bson.D{{Key: "name", Value: name}}
	return uc.GetSinglePersonByParam(filter)
}

func (uc *MongoPerson) GetSinglePersonByParam(filter bson.D) (domain.Person, error) {
	ctx, cancel := context.WithTimeout(context.Background(), uc.ContextTimeout)
	defer cancel()

	var person domain.Person
	err := uc.personRepo.FindOne(ctx, filter).Decode(&person)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.Person{}, errors.New(NOT_FOUND)
		}
		return domain.Person{}, err
	}

	return person, nil
}

func (uc *MongoPerson) Create(person *domain.Person) error {
	ctx, cancel := context.WithTimeout(context.Background(), uc.ContextTimeout)
	defer cancel()

	_, err := uc.personRepo.InsertOne(ctx, person)

	return err
}
