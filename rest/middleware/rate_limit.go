package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	ERROR_TO_CREATE_FOLDER = "failed to create storage directory\n"
	FILE_STORAGE_ERROR     = "File storage error\n"
	LIMIT_EXCEEDED         = "Rate limit exceeded\n"
)

type RateLimitMiddleware struct {
	rateLimit   int
	burst       int
	ttl         time.Duration
	mutex       sync.Mutex
	storagePath string
}

type RateLimitData struct {
	Counter   int64 `json:"counter"`
	Timestamp int64 `json:"timestamp"`
}

func NewRateLimitMiddleware(rateLimit, burst int, ttl time.Duration, storagePath string) *RateLimitMiddleware {
	// Create the storage directory if it doesn't exist
	err := os.MkdirAll(storagePath, os.ModePerm)
	if err != nil {
		panic(ERROR_TO_CREATE_FOLDER + err.Error())
	}
	log.Println("Storage directory:", storagePath)

	return &RateLimitMiddleware{
		rateLimit:   rateLimit,
		burst:       burst,
		ttl:         ttl,
		mutex:       sync.Mutex{},
		storagePath: storagePath,
	}
}

func (m *RateLimitMiddleware) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the client IP address
		clientIP := c.ClientIP()

		// Get the current time in milliseconds
		currentTime := time.Now().UnixNano() / 1000000

		// Key for the rate limit counter
		key := fmt.Sprintf("rate_limit:%s", clientIP)

		// Path to the storage file
		filePath := filepath.Join(m.storagePath, key)

		// Acquire the mutex to prevent concurrent access
		m.mutex.Lock()
		defer m.mutex.Unlock()
		// Check if the file exists
		var rateLimitData RateLimitData
		if _, err := os.Stat(filePath); err == nil {
			// If the file exists, read the rate limit counter and timestamp
			data, err := os.ReadFile(filePath)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": FILE_STORAGE_ERROR + err.Error()})
				return
			}

			if err := json.Unmarshal(data, &rateLimitData); err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": FILE_STORAGE_ERROR + err.Error()})
				return
			}
		} else if !os.IsNotExist(err) {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": FILE_STORAGE_ERROR + err.Error()})
			return
		}

		fmt.Println("DATA: ")
		fmt.Printf("%d - %d\n", rateLimitData.Counter, rateLimitData.Timestamp)
		fmt.Println("OPERATIONS:")
		fmt.Printf("%d - %d < %d\n", currentTime, rateLimitData.Timestamp, m.ttl)

		// Check if the rate limit has been exceeded
		if time.Duration(currentTime-rateLimitData.Timestamp) < m.ttl {
			// Check if the burst limit has been exceeded
			if rateLimitData.Counter >= int64(m.burst) {
				c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": LIMIT_EXCEEDED})
				return
			}
		} else {
			// Reset rate limit data if TTL has expired
			rateLimitData.Counter = 0
			rateLimitData.Timestamp = currentTime
		}

		// Increment the rate limit counter
		rateLimitData.Counter++

		// Update the rate limit data
		data, err := json.Marshal(rateLimitData)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": FILE_STORAGE_ERROR + err.Error()})
			return
		}

		if err := os.WriteFile(filePath, data, 0644); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": FILE_STORAGE_ERROR + err.Error()})
			return
		}

		// Continue processing the request
		c.Next()
	}
}
