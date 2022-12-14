package worker

import (
	"fmt"
	"jeu-de-bourse/internal/models"
	"time"

	"github.com/yyewolf/goronext"
)

func TradesWorker() {
	// Ticker to update the stocks every 15 minutes
	ticker := time.NewTicker(30 * time.Minute)
	go func() {
		for {
			select {
			case <-ticker.C:
				updateTrades()
			}
		}
	}()
	time.Sleep(5 * time.Second)
	updateTrades()
}

func updateTrades() {
	stocks, err := models.GetAllStocks()
	if err != nil {
		fmt.Println("[WORKER-TRADE] Error while getting stocks: ", err)
		return
	}
	fmt.Println("[WORKER-TRADE] Trade update started:", len(stocks))
	errAmount := 0
	for _, s := range stocks {
		stockHistory, err := goronext.GetAllIntraDay(s.ID, s.Market)
		if err != nil {
			fmt.Printf("[WORKER-TRADE] Error while getting stock history: %s\n", err)
			errAmount++
		}

		for _, trade := range stockHistory {
			sh := &models.StockHistory{
				TradeID:   trade.TradeId,
				StockID:   s.ID,
				Price:     trade.Price,
				Time:      trade.Time,
				TradeType: trade.TradeType,
				Volume:    trade.Volume,
			}

			sh.Create()
		}
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("[WORKER-TRADE] Trade update finished:", len(stocks), "stocks updated, ", errAmount, "errors")
}
