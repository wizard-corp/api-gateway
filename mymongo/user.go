package mymongo

import (
	"github.com/wizard-corp/api-gateway/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	db *mongo.Collection
}

func NewUserRepository(mdb MongoDB, collection string) domain.UserRepository {
	return &userRepository{db: mdb.Database.Collection(collection)}
}

func (ur *userRepository) GetByEmail(email string) (domain.User, error) {
	var user domain.User
	err := ur.db.FindOne(c, bson.M{"email": email}).Decode(&user)
	return user, err
}
