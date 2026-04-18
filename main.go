package main

import (
	"coin-market-api/markets"

	"github.com/gin-gonic/gin"
)

func main() {
	InitiateRouter()
}

func InitiateRouter() {
	router := gin.Default()
	api := router.Group("/api/v1")

	markets.Initiate(api)

	router.Run(":8000")
}
