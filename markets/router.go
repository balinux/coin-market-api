package markets

import (
	"net/http"

	"coin-market-api/common/response"

	"github.com/gin-gonic/gin"
)

func Initiate(router *gin.RouterGroup) {
	marketService := NewService()

	markets := router.Group("/markets")
	markets.GET("", func(c *gin.Context) {
		// service
		GetAllMarkets(c, marketService)
	})

	// detail coin
	markets.GET("/:id", func(c *gin.Context) {
		GetCoinDetailHandler(c, marketService)
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

func GetCoinDetailHandler(c *gin.Context, service Service) {
	id := c.Param("id")

	data, err := service.GetCoinDetail(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, response.APIResponse{
		Success: true,
		Message: "Successfully get coin detail",
		Data:    data,
	})
}
