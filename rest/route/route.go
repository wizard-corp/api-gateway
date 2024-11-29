package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wizard-corp/api-gateway/rest/middleware"
	"github.com/wizard-corp/api-gateway/src/bootstrap"
)

func Setup(app *bootstrap.App, timeout time.Duration, gin *gin.Engine) {
	publicRouter := gin.Group("")
	//Middleware log
	publicRouter.Use(middleware.RequestLogger("./log/request_public"))
	// Middleware to limit request
	publicRouter.Use(middleware.NewRateLimitMiddleware(app.Env.RateLimit, app.Env.Burts, time.Duration(app.Env.TTL), "./tmp/rate_limits").Middleware())
	SayHello(app, timeout, publicRouter)
	Login(app, timeout, publicRouter)
	Signup(app, timeout, publicRouter)
	RefreshToken(app, timeout, publicRouter)

	protectedRouter := gin.Group("")
	publicRouter.Use(middleware.RequestLogger("./log/request_private"))
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(app.Env.AccessTokenSecret))
	GetPersonByID(app, timeout, protectedRouter)
	GetPersonByName(app, timeout, protectedRouter)
	CreatePerson(app, timeout, protectedRouter)
}
