package models

import (
	"jeu-de-bourse/internal/database"
	"time"
)

type Stock struct {
	ID       string `json:"id"`
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Exchange string `json:"exchange"`
	Market   string `json:"market"`
	SRD      bool   `json:"srd"`
	Loan     bool   `json:"loan"`
}

// Create a new stock if it doesn't exist
func (s *Stock) Create() error {
	err := database.Session.Query("INSERT INTO jeu_de_bourse.stocks (id, symbol, name, exchange, market, srd, loan) VALUES (?, ?, ?, ?, ?, ?, ?);", s.ID, s.Symbol, s.Name, s.Exchange, s.Market, s.SRD, s.Loan).Exec()
	return err
}

func GetAllStocks() ([]*Stock, error) {
	var stocks []*Stock
	iter := database.Session.Query("SELECT id, symbol, name, exchange, market, srd, loan FROM jeu_de_bourse.stocks;").Iter()
	stock := &Stock{}
	for iter.Scan(&stock.ID, &stock.Symbol, &stock.Name, &stock.Exchange, &stock.Market, &stock.SRD, &stock.Loan) {
		stock = &Stock{}
		stocks = append(stocks, stock)
	}
	return stocks, iter.Close()
}

type StockHistory struct {
	TradeID   string     `json:"trade_id"`
	StockID   string     `json:"stock_id"`
	Price     int        `json:"price"`
	Time      *time.Time `json:"time"`
	TradeType string     `json:"trade_type"`
	Volume    float64    `json:"volume"`
}

// Create a new stock history entry
func (s *StockHistory) Create() error {
	err := database.Session.Query("INSERT INTO jeu_de_bourse.stocks_history (trade_id, stock_id, price, time, trade_type, volume) VALUES (?, ?, ?, ?, ?, ?);", s.TradeID, s.StockID, s.Price, s.Time, s.TradeType, s.Volume).Exec()
	return err
}
