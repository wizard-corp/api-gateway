package main

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/wizard-corp/api-gateway/rest/route"
	"github.com/wizard-corp/api-gateway/src/bootstrap"
	"github.com/wizard-corp/api-gateway/src/mylogger"
)

func main() {
	app := bootstrap.NewApp()

	err := mylogger.SetupLogger("./log/app.log", "info") // Set log file and level
	if err != nil {
		panic(err)
	}

	gin := gin.Default()
	route.Setup(app, time.Duration(app.Env.ContextTimeout)*time.Second, gin)
	gin.Run(app.Env.ServerAddress) // listen and serve on 0.0.0.0:8080
}
