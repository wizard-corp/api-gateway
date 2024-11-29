package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wizard-corp/api-gateway/src/domain"
)

// ErrorHandlerMiddleware handles errors based on their type and sends appropriate HTTP responses
func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Process request

		// After request processing, check if any errors occurred
		if len(c.Errors) > 0 {
			for _, ginErr := range c.Errors {
				switch err := ginErr.Err.(type) {
				case *domain.PresentationError:
					// Presentation errors are returned with a 400 Bad Request
					c.JSON(http.StatusBadRequest, gin.H{
						"error": err.Error(),
					})
					log.Printf("PresentationError: %v", err.Error())

				case *domain.InfrastructureError:
					// Infrastructure errors are returned with a 500 Internal Server Error
					c.JSON(http.StatusInternalServerError, gin.H{
						"error": "Internal Server Error",
					})
					log.Printf("InfrastructureError: %v", err.Error())

				case *domain.ApplicationError:
					// Application errors are returned with a 422 Unprocessable Entity
					c.JSON(http.StatusUnprocessableEntity, gin.H{
						"error": "Application Error: Logic Failure",
					})
					log.Printf("ApplicationError: %v", err.Error())

				case *domain.DomainError:
					// Domain errors are returned with a 400 Bad Request
					c.JSON(http.StatusBadRequest, gin.H{
						"error": err.Error(),
					})
					log.Printf("DomainError: %v", err.Error())

				default:
					// Unknown or unhandled errors
					c.JSON(http.StatusInternalServerError, gin.H{
						"error": "Unknown error occurred",
					})
					log.Printf("Unknown error: %v", err)
				}
			}
			// Stop further handlers from running
			c.Abort()
		}
	}
}
