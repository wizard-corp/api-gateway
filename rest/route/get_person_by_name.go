package route

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wizard-corp/api-gateway/src/bootstrap"
	"github.com/wizard-corp/api-gateway/src/presentation"
)

func GetPersonByName(app *bootstrap.App, timeout time.Duration, group *gin.RouterGroup) {
	fn := func(c *gin.Context) {
		personId := strings.TrimSpace(c.Param("name"))
		if personId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid identifier"})
			return
		}

		lc := presentation.NewPersonController(timeout, app)
		response, err := lc.GetPersonByName(personId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, response)
	}
	group.GET("/person?name", fn)
}
