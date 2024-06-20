package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wizard-corp/api-gateway/bootstrap"
)

func SayHello(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	lc := func(c *gin.Context) {
		c.JSON(
			200,
			gin.H{"message": "Hello!"})
	}
	group.GET("hello", lc)
}
