package route

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wizard-corp/api-gateway/src/bootstrap"
	"github.com/wizard-corp/api-gateway/src/domain"
	"github.com/wizard-corp/api-gateway/src/presentation"
)

type NewPersonRequest struct {
	FirtsName   string `form:"firtsName" binding:"required"`
	PatriLineal string `form:"patriLineal" binding:"required"`
	MatriLineal string `form:"matriLineal" binding:"required"`
	Address     string `form:"address" binding:"required"`
	BirthDate   string `form:"birthDate" binding:"required"`
}

func NewPerson(app *bootstrap.App, timeout time.Duration, group *gin.RouterGroup) {
	fn := func(c *gin.Context) {
		var request NewPersonRequest
		err := c.ShouldBind(&request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": domain.INVALID_SCHEMA + "\n" + err.Error()})
			return
		}

		lc := presentation.NewPersonController(timeout, app)
		err = lc.NewPerson(request.FirtsName, request.PatriLineal, request.MatriLineal, request.Address, request.BirthDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
	group.POST("/person", fn)
}
