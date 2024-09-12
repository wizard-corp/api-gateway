package main

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/wizard-corp/api-gateway/rest/route"
	"github.com/wizard-corp/api-gateway/src/bootstrap"
)

func main() {
	app := bootstrap.NewApp()

	timeout := time.Duration(app.Env.ContextTimeout) * time.Second

	gin := gin.Default()
	route.Setup(app, timeout, gin)
	gin.Run(app.Env.ServerAddress) // listen and serve on 0.0.0.0:8080
}
