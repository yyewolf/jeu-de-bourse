package stocks

import (
	"fmt"
	"jeu-de-bourse/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// GetParams
type GetParams struct {
	// Page
	Page int `json:"page" binding:"required"`
}

const (
	// StocksPerPage
	StocksPerPage = 20
)

type EnrichedStock struct {
	*models.Stock
	OpeningPrice int `json:"opening_price"`
	LastPrice    int `json:"last_price"`
}

func GetStocksPage(c *gin.Context) {
	var params GetParams
	if err := c.ShouldBindBodyWith(&params, binding.JSON); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"status":  "Les informations fournies ne sont pas valides",
		})
		return
	}

	s, err := models.GetAllStocksPaged(params.Page, StocksPerPage)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"success": false,
			"status":  "Une erreur est survenue",
		})
		return
	}

	a, err := models.GetStocksAmountOfPages(StocksPerPage)
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"status":  "Une erreur est survenue",
		})
		return
	}

	retStocks := make([]*EnrichedStock, len(s))

	for i, stock := range s {
		eStock := &EnrichedStock{
			Stock:        stock,
			OpeningPrice: 0,
			LastPrice:    0,
		}

		retStocks[i] = eStock

		prices, err := stock.GetStockValuesToday()
		if err != nil {
			continue
		}

		if len(prices) > 0 {
			eStock.LastPrice = prices[0]
			eStock.OpeningPrice = prices[len(prices)-1]
		}
	}

	c.JSON(200, gin.H{
		"success": true,
		"stocks":  retStocks,
		"pages":   a,
	})
}
