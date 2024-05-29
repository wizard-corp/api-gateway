package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/wizard-corp/api-gateway/bootstrap"
	"github.com/wizard-corp/api-gateway/mymongo"
)

func main() {
	app := bootstrap.App()
	env := app.Env

	server := &mymongo.MongoServer{
		Host:     "localhost",
		Port:     27017,
		User:     "mongo",
		Password: "mongo",
		Database: "test",
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		fmt.Println("Received a GET request to the root path")
		err := mymongo.TestInfrastructure(server)
		c.JSON(200, gin.H{
			"message": err,
		})
	})
	r.Run(env.ServerAddress) // listen and serve on 0.0.0.0:8080
}
