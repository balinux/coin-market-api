package markets

import (
	"net/http"

	"go-mrt-schedule/common/response"

	"github.com/gin-gonic/gin"
)

func Initiate(router *gin.RouterGroup) {
	marketService := NewService()

	markets := router.Group("/markets")
	markets.GET("", func(c *gin.Context) {
		// service
		GetAllMarkets(c, marketService)
	})
}

func GetAllMarkets(c *gin.Context, service Service) {
	data, err := service.GetAllMarkets()
	if err != nil {
		// handler error
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	// handle response
	c.JSON(http.StatusOK, response.APIResponse{
		Success: true,
		Message: "Successfully get all markets",
		Data:    data,
	})
}
