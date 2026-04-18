package markets

import "github.com/gin-gonic/gin"

func Initiate(router *gin.RouterGroup) {
	marketService := NewService()

	markets := router.Group("/markets")
	markets.GET("", func(c *gin.Context) {
		// service
		GetAllMarkets(c, marketService)
	})
}

func GetAllMarkets(c *gin.Context, service Service) {
	data, err := service.GetAllMarkets()()
	if err != nil {
		// handler error
	}
	// handle response
}
