package models

import (
	"fmt"
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

func GetAllStocksPaged(page int, limit int) ([]*Stock, error) {
	// We are using CQL so there is no OFFSET
	stocks, err := GetAllStocks()
	if err != nil {
		return nil, err
	}
	if page*limit > len(stocks) {
		return nil, fmt.Errorf("page %d is out of range", page)
	}
	return stocks[page*limit : (page+1)*limit], nil
}

func GetStocksAmountOfPages(limit int) (int, error) {
	var count int
	err := database.Session.Query("SELECT COUNT(*) FROM jeu_de_bourse.stocks;").Scan(&count)
	return count / limit, err
}

func FindStockWithQuery(query string) ([]*Stock, error) {
	var stocks []*Stock
	iter := database.Session.Query("SELECT id, symbol, name, exchange, market, srd, loan FROM jeu_de_bourse.stocks WHERE symbol LIKE ?;", "%"+query+"%").Iter()
	stock := &Stock{}
	for iter.Scan(&stock.ID, &stock.Symbol, &stock.Name, &stock.Exchange, &stock.Market, &stock.SRD, &stock.Loan) {
		stock = &Stock{}
		stocks = append(stocks, stock)
	}
	return stocks, iter.Close()
}

// Get stock value at a specific time
func (s *Stock) GetStockValueAtTime(t time.Time) (int, error) {
	var price int
	err := database.Session.Query("SELECT price FROM jeu_de_bourse.stocks_history WHERE stock_id = ? AND time <= ? LIMIT 1;", s.ID, t).Scan(&price)
	return price, err
}

// Get stock overall today value
func (s *Stock) GetStockValuesToday() ([]int, error) {
	var prices []int
	// find today's midnight
	midnight := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Now().Location())
	iter := database.Session.Query("SELECT price FROM jeu_de_bourse.stocks_history WHERE stock_id = ? AND time >= ?;", s.ID, midnight).Iter()
	var price int
	for iter.Scan(&price) {
		prices = append(prices, price)
	}
	return prices, iter.Close()
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
