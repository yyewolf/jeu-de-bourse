package worker

import (
	"fmt"
	"jeu-de-bourse/internal/models"
	"time"

	"github.com/yyewolf/goronext"
)

func TradesWorker() {
	// Ticker to update the stocks every 15 minutes
	ticker := time.NewTicker(10 * time.Minute)
	go func() {
		for {
			select {
			case <-ticker.C:
				updateTrades("100")
				// At midnight, update slow to make sure we have all the data
				if time.Now().Hour() == 0 {
					updateTrades("50000")
				}
			}
		}
	}()
	time.Sleep(5 * time.Second)
	updateTrades("100")
}

func updateTrades(amount string) {
	stocks, err := models.GetAllStocks()
	if err != nil {
		fmt.Println("[WORKER-TRADE] Error while getting stocks: ", err)
		return
	}
	fmt.Println("[WORKER-TRADE] Trade update started:", len(stocks))
	errAmount := 0
	for _, s := range stocks {
		when := time.Now()
		// If it's saturday or sunday, we look at friday's data
		if when.Weekday() == time.Saturday {
			when = when.AddDate(0, 0, -1)
		} else if when.Weekday() == time.Sunday {
			when = when.AddDate(0, 0, -2)
		}

		stockHistory, err := goronext.GetAllIntraDay(s.ID, s.Market, amount, when)
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
