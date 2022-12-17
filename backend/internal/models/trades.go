package models

import (
	"jeu-de-bourse/internal/database"
	"time"
)

const (
	TradeStateCreated   TradeState = "created"
	TradeStateStarted   TradeState = "started"
	TradeStateTryToStop TradeState = "stopping"
	TradeStateStopped   TradeState = "stopped"
)

type TradeState string

type Trade struct {
	// Partition key
	UserID string `json:"user_id"`
	// Clustering keys
	State   TradeState `json:"state"`
	StockID string     `json:"stock_id"`
	SRD     bool       `json:"srd"`
	ID      string     `json:"id"`

	Amount     int `json:"amount"`
	StartPrice int `json:"start_price"`
	StopPrice  int `json:"stop_price"`
	Threshold  int `json:"threshold"`
	Limit      int `json:"lim"`

	CreatedAt   time.Time `json:"created_at"`
	StartedAt   time.Time `json:"started_at"`
	TryToStopAt time.Time `json:"try_to_stop_at"`
	StoppedAt   time.Time `json:"stopped_at"` // Clustering order by this field
}

func GetCompletedTradesByUserID(userID string) ([]Trade, error) {
	var trades []Trade

	if err := database.Session.Query("SELECT * FROM jeu_de_bourse.trades WHERE user_id = ? AND state = ? ALLOW FILTERING;", userID, TradeStateStopped).Scan(&trades); err != nil {
		return nil, err
	}

	return trades, nil
}

func GetStartedTradesByUserID(userID string) ([]Trade, error) {
	var trades []Trade

	if err := database.Session.Query("SELECT * FROM jeu_de_bourse.trades WHERE user_id = ? AND state = ? ALLOW FILTERING;", userID, TradeStateStarted).Scan(&trades); err != nil {
		return nil, err
	}

	return trades, nil
}

func GetCreatedTradesByUserID(userID string) ([]Trade, error) {
	var trades []Trade

	if err := database.Session.Query("SELECT * FROM jeu_de_bourse.trades WHERE user_id = ? AND state = ? ALLOW FILTERING;", userID, TradeStateCreated).Scan(&trades); err != nil {
		return nil, err
	}

	return trades, nil
}

func GetStoppingTradesByUserID(userID string) ([]Trade, error) {
	var trades []Trade

	if err := database.Session.Query("SELECT * FROM jeu_de_bourse.trades WHERE user_id = ? AND state = ? ALLOW FILTERING;", userID, TradeStateTryToStop).Scan(&trades); err != nil {
		return nil, err
	}

	return trades, nil
}
