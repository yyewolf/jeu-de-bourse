package worker

import (
	"fmt"
	"jeu-de-bourse/internal/models"
	"time"

	"github.com/yyewolf/goronext"
)

// Should run once at startup

func LoadStocks() {
	stocks, err := goronext.FindAllStock()
	if err != nil {
		panic(err)
	}

	for _, stock := range stocks {
		companyInfo, err := goronext.GetCompanyInfo(stock.ID, stock.Market)
		if err != nil {
			fmt.Printf("[WORKER-SCAN] Error while getting company info: %s\n", err)
		}

		s := &models.Stock{
			ID:       stock.ID,
			Symbol:   stock.Symbol,
			Name:     stock.Name,
			Exchange: stock.Exchange,
			Market:   stock.Market,
			SRD:      companyInfo.SRD,
			Loan:     companyInfo.Loan,
		}

		s.Create()
		// fmt.Println("[WORKER-SCAN] Stock created: ", s.Symbol)
		time.Sleep(100 * time.Millisecond)
	}
}
