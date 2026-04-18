package markets

import "github.com/gin-gonic/gin"

func Initiate(router *gin.RouterGroup) {
	markets := router.Group("/markets")
	markets.GET("", func(c *gin.Context) {
		// service
	})
}
