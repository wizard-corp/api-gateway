package route

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wizard-corp/api-gateway/src/bootstrap"
	"github.com/wizard-corp/api-gateway/src/domain"
	"github.com/wizard-corp/api-gateway/src/presentation"
)

type RefreshTokenRequest struct {
	RefreshToken string `form:"refreshToken" binding:"required"`
}

func RefreshToken(app *bootstrap.App, timeout time.Duration, group *gin.RouterGroup) {
	fn := func(c *gin.Context) {
		var request RefreshTokenRequest
		err := c.ShouldBind(&request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": domain.INVALID_SCHEMA + "\n" + err.Error()})
			return
		}

		lc := presentation.NewRefreshTokenController(timeout, app)
		response, err := lc.RefreshToken(request.RefreshToken, app.Env.AccessTokenSecret, app.Env.AccessTokenExpiryHour, app.Env.RefreshTokenSecret, app.Env.RefreshTokenExpiryHour)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, response)
	}
	group.POST("/refresh", fn)
}
