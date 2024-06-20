package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/wizard-corp/api-gateway/api/route"
	"github.com/wizard-corp/api-gateway/bootstrap"
	"github.com/wizard-corp/api-gateway/mymongo"
)

// TODO
// create package rest/route
// create handler for route ping
func main() {
	app := bootstrap.NewApp()
	env := app.Env

	timeout := time.Duration(env.ContextTimeout) * time.Second

	fmt.Println("Test Infrastructure...")
	err := mymongo.TestInfrastructure(app.Mongo)
	if err != nil {
		fmt.Println(err)
	}

	gin := gin.Default()
	route.Setup(app, timeout, gin)
	gin.Run(env.ServerAddress) // listen and serve on 0.0.0.0:8080
}
