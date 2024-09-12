package bootstrap

import (
	"log"

	"github.com/wizard-corp/api-gateway/src/mymongo"
)

type App struct {
	Env   *Env
	Mongo *mymongo.MongoDB
}

func NewApp() *App {
	env := NewEnv()

	mongoDBClient, err := mymongo.NewMongoDBClient(&mymongo.MongoConfig{
		Host:     env.MongoHost,
		Port:     env.MongoPort,
		User:     env.MongoUser,
		Password: env.MongoPassword,
		Database: env.MongoDatabase,
	})
	if err != nil {
		log.Fatal(err)
	}

	return &App{Env: env, Mongo: mongoDBClient}
}

func (app *App) CloseApp() error {
	app.Mongo.Close()

	return nil
}
