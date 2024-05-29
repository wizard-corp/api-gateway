package bootstrap

import (
	"github.com/wizard-corp/api-gateway/mymongo"
)

type Application struct {
	Env   *Env
	Mongo *mymongo.MongoServer
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Mongo = &mymongo.MongoServer{
		Host:     app.Env.MongoHost,
		Port:     app.Env.MongoPort,
		User:     app.Env.MongoUser,
		Password: app.Env.MongoPassword,
		Database: app.Env.MongoDatabase,
	}
	return *app
}
