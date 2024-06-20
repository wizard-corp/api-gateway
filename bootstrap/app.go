package bootstrap

import (
	"github.com/wizard-corp/api-gateway/mymongo"
)

type App struct {
	Env   *Env
	Mongo *mymongo.MongoServer
}

func NewApp() App {
	app := &App{}
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
