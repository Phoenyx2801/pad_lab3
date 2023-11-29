package main

import (
	contr "REST-api/pkg/controlers"
	"REST-api/pkg/db"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
	"os"
	"time"
)

var redisClient *redis.Client

func initRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: "",
		DB:       0,
	})
}

func checkRedisConnection() error {
	_, err := redisClient.Ping(context.Background()).Result()
	return err
}

func getBooksFromCacheOrDB(c *gin.Context) {
	// Check Redis connection (initialize only once, not on every request)
	if redisClient == nil {
		initRedis()
		err := checkRedisConnection()
		if err != nil {
			fmt.Println("Error connecting to Redis:", err)
			// Proceed to fetch data from the database
			contr.GetBooks(c)
			return
		}
	}

	// Define a key for caching based on the request URL
	cacheKey := c.Request.URL.String()
	fmt.Println("Cache key: " + cacheKey)

	// Check if the data is available in the cache
	cachedData, err := redisClient.Get(context.Background(), cacheKey).Result()
	if err == nil {
		// Data found in cache, return it as JSON
		fmt.Println("Data found in cache")

		// Return the cached data to the client
		c.JSON(http.StatusOK, gin.H{"data": cachedData})
		return
	}

	// Data not found in cache, fetch it from the database
	fmt.Println("Data not found in cache. Fetching from the database.")

	// Call the original controller function to get the data from the database
	contr.GetBooks(c)

	// Get the response body
	responseBody, exists := c.Get("responseBody")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error 1"})
		return
	}

	// Convert the response body to string
	var responseBodyString string
	switch responseBody := responseBody.(type) {
	case string:
		responseBodyString = responseBody
	default:
		// If responseBody is not a string, try to marshal it to JSON
		jsonBody, err := json.Marshal(responseBody)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error 2"})
			return
		}
		responseBodyString = string(jsonBody)
	}

	// Store the data in the cache with the key being the request URL and a cache time of 3 minutes
	err = redisClient.Set(context.Background(), cacheKey, responseBodyString, 3*time.Minute).Err()
	if err != nil {
		fmt.Println("Error caching data:", err)
	}

	// Return the data to the client in JSON format
	//c.JSON(http.StatusOK, gin.H{"data": responseBodyString})
}

func main() {

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/search/:id", getBooksFromCacheOrDB)
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, "pong")
	})
	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "App run")
	})
	r.POST("/addBook", contr.AddBook)
	r.DELETE("/delete/:id", contr.DeleteBook)
	r.PUT("/update/:id", contr.UpdateBookById)

	db.ConnectPostgres()
	err := r.Run(":" + os.Getenv("APP_PORT"))
	if err != nil {
		return
	}
}
