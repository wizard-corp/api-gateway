package bootstrap

import (
	"log"

	"github.com/wizard-corp/api-gateway/domain"
	"github.com/wizard-corp/api-gateway/mymongo"
)

type App struct {
	Env   *Env
	Mongo *mymongo.MongoDB
}

func NewApp() App {
	app := &App{}
	app.Env = NewEnv()
	mongoClient, err := mymongo.NewMongoClient(&mymongo.MongoServer{
        Host:     app.Env.MongoHost,
        Port:     app.Env.MongoPort,
        User:     app.Env.MongoUser,
        Password: app.Env.MongoPassword,
        Database: app.Env.MongoDatabase,
    })
    if err != nil {
        return nil, err
    }
    app.Mongo = mongoClient

    return app, nil
}
