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
	FirtsName   string `json:"firtsName" binding:"required"`
	PatriLineal string `json:"patriLineal" binding:"required"`
	MatriLineal string `json:"matriLineal" binding:"required"`
	Address     string `json:"address" binding:"required"`
	BirthDate   string `json:"birthDate" binding:"required"`
}

func CreatePerson(app *bootstrap.App, timeout time.Duration, group *gin.RouterGroup) {
	fn := func(c *gin.Context) {
		var request NewPersonRequest
		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": domain.INVALID_SCHEMA + "\n" + err.Error()})
			return
		}

		lc := presentation.NewPersonController(timeout, app)
		err := lc.CreatePerson(request.FirtsName, request.PatriLineal, request.MatriLineal, request.Address, request.BirthDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
	group.POST("/person", fn)
}
