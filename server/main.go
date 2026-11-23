package main

import (
	"fmt"

	"github.com/MaitreyeeBudhale/rate-limiter.git/generators"
	"github.com/MaitreyeeBudhale/rate-limiter.git/limiters"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "hello"})
	})

	limiter := limiters.NewTokenBucket(
		10, // capacity
		5,  // refill rate
		1,  // refill period
	)

	stats := generators.RateLimiter(
		100, // total requests
		10,  // concurrent workers
		limiter,
	)

	fmt.Println("Total Requests:", 100)
	fmt.Println("Accepted:", stats.Accepted)
	fmt.Println("Rejected:", stats.Rejected)

	router.Run(":3000")
}
