package mymongo

import (
	"context"
	"errors"
	"time"

	"github.com/wizard-corp/api-gateway/src/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoUser struct {
	userRepo       *mongo.Collection
	ContextTimeout time.Duration
}

func isEmptyUser(user domain.User) bool {
	return user.Email == ""
}

func NewUserRepository(mongodb *MongoDB, timeout time.Duration) *MongoUser {
	collection := mongodb.Mdb.Collection("user")
	return &MongoUser{collection, timeout}
}

func (uc *MongoUser) Fetch() ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), uc.ContextTimeout)
	defer cancel()

	var data []domain.User
	findOptions := options.Find()
	//findOptions.SetLimit(5)
	cur, err := uc.userRepo.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		return data, errors.New(DB_CURSOR_FAILED)
	}

	for cur.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var elem domain.User
		err := cur.Decode(&elem)
		if err != nil {
			return data, errors.New(DECODE_FAILED)
		}

		data = append(data, elem)

	}
	if err := cur.Err(); err != nil {
		return data, errors.New(DB_CURSOR_FAILED)
	}
	cur.Close(ctx)
	//fmt.Printf("Found multiple documents: %+v\n", data)

	return data, nil
}

func (uc *MongoUser) GetUserByEmail(email string) (domain.User, error) {
	filter := bson.D{{Key: "email", Value: email}}
	return uc.GetSingleUserByParam(filter)
}

func (uc *MongoUser) GetUserByID(id string) (domain.User, error) {
	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.User{}, err
	}
	filter := bson.D{{Key: "_id", Value: idHex}}
	return uc.GetSingleUserByParam(filter)
}

func (uc *MongoUser) GetSingleUserByParam(filter bson.D) (domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), uc.ContextTimeout)
	defer cancel()

	var user domain.User
	err := uc.userRepo.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.User{}, errors.New(NOT_FOUND)
		}
		return domain.User{}, err
	}

	if isEmptyUser(user) {
		return domain.User{}, errors.New(NOT_FOUND)
	}

	return user, nil
}

func (uc *MongoUser) Create(user *domain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), uc.ContextTimeout)
	defer cancel()

	user.MongoSanitizeUser()

	_, err := uc.userRepo.InsertOne(ctx, user)

	return err
}
