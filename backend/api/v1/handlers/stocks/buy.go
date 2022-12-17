package stocks

import (
	"jeu-de-bourse/internal/models"

	"github.com/gin-gonic/gin"
)

type buyStockForm struct {
	ID        string `form:"id" json:"id" binding:"required"`
	Amount    int    `form:"amount" json:"amount" binding:"required"`
	Limit     int    `form:"limit" json:"limit" binding:"required"`
	Threshold int    `form:"threshold" json:"threshold" binding:"required"`
}

func BuyStock(c *gin.Context) {
	var form buyStockForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(400, gin.H{
			"success": false,
		})
		return
	}

	stock, err := models.GetStockByID(form.ID)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"status":  "stock not found",
		})
		return
	}

	// If amount is negative, it's a loan so we must check if the stock is loanable
	if form.Amount < 0 && !stock.Loan {
		c.JSON(400, gin.H{
			"success": false,
			"status":  "stock is not loanable",
		})
		return
	}

	// Can only be limiting or thresholding, not both
	if form.Limit != -666666 && form.Threshold != -666666 {
		c.JSON(400, gin.H{
			"success": false,
			"status":  "can only be limiting or thresholding, not both",
		})
		return
	}

}
