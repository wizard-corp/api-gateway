package route

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wizard-corp/api-gateway/src/bootstrap"
	"github.com/wizard-corp/api-gateway/src/domain"
	"github.com/wizard-corp/api-gateway/src/presentation"
)

type LoginRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

func Login(app *bootstrap.App, timeout time.Duration, group *gin.RouterGroup) {
	fn := func(c *gin.Context) {
		var request LoginRequest
		err := c.ShouldBind(&request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": domain.INVALID_SCHEMA + "\n" + err.Error()})
			return
		}

		lc := presentation.NewLoginController(timeout, app)
		response, err := lc.Login(request.Email, request.Password, app.Env.AccessTokenSecret, app.Env.AccessTokenExpiryHour, app.Env.RefreshTokenSecret, app.Env.RefreshTokenExpiryHour)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, response)
	}
	group.POST("/login", fn)
}
