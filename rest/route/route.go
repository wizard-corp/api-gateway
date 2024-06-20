package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wizard-corp/api-gateway/bootstrap"
)

func Setup(app *bootstrap.App, timeout time.Duration, gin *gin.Engine) {
	publicRouter := gin.Group("")
	// All Public APIs
	SayHello(app, timeout, publicRouter)
	Login(app, timeout, publicRouter)
}
