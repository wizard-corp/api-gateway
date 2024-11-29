package middleware

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CustomResponseWriter to capture the response body
type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b) // Capture the response body
	return w.ResponseWriter.Write(b)
}

// LoggerMiddleware logs request and response details
func RequestLogger(filename string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Request start time
		requestStart := time.Now()

		// Generate a unique request ID
		requestID := uuid.New().String()

		// Extract necessary information from the request
		region := c.GetHeader("Region")
		protocol := c.Request.Proto
		ipAddress := c.ClientIP()
		requestPath := c.Request.URL.Path
		requestHeaders := c.Request.Header
		requestMethod := c.Request.Method
		requestParams := c.Request.URL.RawQuery
		requestSize := c.Request.ContentLength

		// Read request body size
		var requestBodySize int64
		if c.Request.Body != nil {
			bodyBytes, _ := io.ReadAll(c.Request.Body)
			requestBodySize = int64(len(bodyBytes))
			// Restore request body for further processing
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		// Use custom response writer to capture response
		customWriter := &CustomResponseWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = customWriter

		// Proceed with the request
		c.Next()

		// Request end time
		requestEnd := time.Now()

		// Extract response information
		status := customWriter.Status()
		responseBody := customWriter.body.String()
		responseHeaders := customWriter.Header()
		responseSize := customWriter.Size()

		// Log the details into a log file
		logData := fmt.Sprintf(
			"RequestID: %s\nRegion: %s\nRequestStart: %s\nRequestEnd: %s\nProtocol: %s\nIPAddress: %s\nRequestPath: %s\nRequestHeaders: %v\nRequestMethod: %s\nRequestParams: %s\nRequestSize: %d\nRequestBodySize: %d\nStatus: %d\nResponseBody: %s\nResponseHeaders: %v\nResponseSize: %d\n\n",
			requestID, region, requestStart.Format(time.RFC3339), requestEnd.Format(time.RFC3339), protocol,
			ipAddress, requestPath, requestHeaders, requestMethod, requestParams, requestSize,
			requestBodySize, status, responseBody, responseHeaders, responseSize,
		)

		// Open the log file in append mode, create if it doesn't exist
		logFile, err := os.OpenFile(filename+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("Failed to open log file: %v", err)
		}
		defer logFile.Close()

		// Write the log data into the file
		if _, err := logFile.WriteString(logData); err != nil {
			log.Fatalf("Failed to write log: %v", err)
		}
	}
}
