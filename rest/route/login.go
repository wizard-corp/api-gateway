package route

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wizard-corp/api-gateway/bootstrap"
	"github.com/wizard-corp/api-gateway/domain"
	"github.com/wizard-corp/api-gateway/presentation"
)

type LoginRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

func Login(app *bootstrap.App, timeout time.Duration, group *gin.RouterGroup) {
	lc := func(c *gin.Context) {
		var request LoginRequest
		err := c.ShouldBind(&request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": domain.INVALID_SCHEMA})
			return
		}

		loginUseCase := domain.NewLoginUsecase(app.Env.SystemUId, app.Mongo)
		response, err := presentation.NewLoginController(
			loginUseCase,
			request.Email,
			request.Password,
			app.Env.AccessTokenSecret,
			app.Env.AccessTokenExpiryHour,
			app.Env.RefreshTokenSecret,
			app.Env.RefreshTokenExpiryHour,
		)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": domain.LOGIC_CRUSH})
			return
		}

		c.JSON(http.StatusOK, response)
	}
	group.POST("/login", lc)
}
