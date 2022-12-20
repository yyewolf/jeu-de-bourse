package models

import (
	"jeu-de-bourse/internal/database"
	"time"

	"github.com/gocql/gocql"
)

const (
	TradeStateCreated   TradeState = "created"
	TradeStateStarted   TradeState = "started"
	TradeStateTryToStop TradeState = "stopping"
	TradeStateStopped   TradeState = "stopped"
)

type TradeState string

// Implement marshal for TradeState
func (t TradeState) MarshalCQL(info gocql.TypeInfo) ([]byte, error) {
	return []byte(t), nil
}

func (t *TradeState) UnmarshalCQL(info gocql.TypeInfo, data []byte) error {
	*t = TradeState(data)
	return nil
}

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

func (t *Trade) Create() error {
	if err := database.Session.Query("INSERT INTO jeu_de_bourse.trades (user_id, state, stock_id, srd, id, amount, start_price, stop_price, threshold, lim, created_at, started_at, try_to_stop_at, stopped_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);", t.UserID, t.State, t.StockID, t.SRD, t.ID, t.Amount, t.StartPrice, t.StopPrice, t.Threshold, t.Limit, t.CreatedAt, t.StartedAt, t.TryToStopAt, t.StoppedAt).Exec(); err != nil {
		return err
	}
	return nil
}

func (t *Trade) Update() error {
	if err := database.Session.Query("UPDATE jeu_de_bourse.trades SET state = ?, amount = ?, start_price = ?, stop_price = ?, threshold = ?, lim = ?, created_at = ?, started_at = ?, try_to_stop_at = ?, stopped_at = ? WHERE user_id = ? AND state = ? AND stock_id = ? AND srd = ? AND id = ?;", t.State, t.Amount, t.StartPrice, t.StopPrice, t.Threshold, t.Limit, t.CreatedAt, t.StartedAt, t.TryToStopAt, t.StoppedAt, t.UserID, t.State, t.StockID, t.SRD, t.ID).Exec(); err != nil {
		return err
	}
	return nil
}

func GetCompletedTradesByUserID(userID string, srd bool) ([]*Trade, error) {
	var trades []*Trade

	if err := database.Session.Query("SELECT user_id, state, stock_id, srd, id, amount, start_price, stop_price, threshold, lim, created_at, started_at, try_to_stop_at, stopped_at FROM jeu_de_bourse.trades WHERE user_id = ? AND state = ? AND srd = ? ALLOW FILTERING;", userID, TradeStateStopped, srd).Scan(&trades); err != nil {
		return nil, err
	}

	return trades, nil
}

func GetStartedTradesByUserID(userID string, srd bool) ([]*Trade, error) {
	var trades []*Trade

	iter := database.Session.Query("SELECT user_id, state, stock_id, srd, id, amount, start_price, stop_price, threshold, lim, created_at, started_at, try_to_stop_at, stopped_at FROM jeu_de_bourse.trades WHERE user_id = ? AND state = ? AND srd = ? ALLOW FILTERING;", userID, TradeStateStarted, srd).Iter()
	for {
		var trade Trade
		if !iter.Scan(&trade.UserID, &trade.State, &trade.StockID, &trade.SRD, &trade.ID, &trade.Amount, &trade.StartPrice, &trade.StopPrice, &trade.Threshold, &trade.Limit, &trade.CreatedAt, &trade.StartedAt, &trade.TryToStopAt, &trade.StoppedAt) {
			break
		}
		trades = append(trades, &trade)
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}

	return trades, nil
}

func GetCreatedTradesByUserID(userID string, srd bool) ([]*Trade, error) {
	var trades []*Trade

	iter := database.Session.Query("SELECT user_id, state, stock_id, srd, id, amount, start_price, stop_price, threshold, lim, created_at, started_at, try_to_stop_at, stopped_at FROM jeu_de_bourse.trades WHERE user_id = ? AND state = ? AND srd = ? ALLOW FILTERING;", userID, TradeStateCreated, srd).Iter()
	for {
		var trade Trade
		if !iter.Scan(&trade.UserID, &trade.State, &trade.StockID, &trade.SRD, &trade.ID, &trade.Amount, &trade.StartPrice, &trade.StopPrice, &trade.Threshold, &trade.Limit, &trade.CreatedAt, &trade.StartedAt, &trade.TryToStopAt, &trade.StoppedAt) {
			break
		}
		trades = append(trades, &trade)
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}

	return trades, nil
}

func GetStoppingTradesByUserID(userID string, srd bool) ([]*Trade, error) {
	var trades []*Trade

	iter := database.Session.Query("SELECT user_id, state, stock_id, srd, id, amount, start_price, stop_price, threshold, lim, created_at, started_at, try_to_stop_at, stopped_at FROM jeu_de_bourse.trades WHERE user_id = ? AND state = ? AND srd = ? ALLOW FILTERING;", userID, TradeStateTryToStop, srd).Iter()
	for {
		var trade Trade
		if !iter.Scan(&trade.UserID, &trade.State, &trade.StockID, &trade.SRD, &trade.ID, &trade.Amount, &trade.StartPrice, &trade.StopPrice, &trade.Threshold, &trade.Limit, &trade.CreatedAt, &trade.StartedAt, &trade.TryToStopAt, &trade.StoppedAt) {
			break
		}
		trades = append(trades, &trade)
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}

	return trades, nil
}

func GetOnGoingTradesByUserIDWithSRD(userID string, srd bool) ([]*Trade, error) {
	var trades []*Trade

	// using 3 iters for each state
	iter := database.Session.Query("SELECT user_id, state, stock_id, srd, id, amount, start_price, stop_price, threshold, lim, created_at, started_at, try_to_stop_at, stopped_at FROM jeu_de_bourse.trades WHERE user_id = ? AND state = ? AND srd = ? ALLOW FILTERING;", userID, TradeStateCreated, srd).Iter()
	for {
		var trade Trade
		if !iter.Scan(&trade.UserID, &trade.State, &trade.StockID, &trade.SRD, &trade.ID, &trade.Amount, &trade.StartPrice, &trade.StopPrice, &trade.Threshold, &trade.Limit, &trade.CreatedAt, &trade.StartedAt, &trade.TryToStopAt, &trade.StoppedAt) {
			break
		}
		trades = append(trades, &trade)
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}

	iter = database.Session.Query("SELECT user_id, state, stock_id, srd, id, amount, start_price, stop_price, threshold, lim, created_at, started_at, try_to_stop_at, stopped_at FROM jeu_de_bourse.trades WHERE user_id = ? AND state = ? AND srd = ? ALLOW FILTERING;", userID, TradeStateStarted, srd).Iter()
	for {
		var trade Trade
		if !iter.Scan(&trade.UserID, &trade.State, &trade.StockID, &trade.SRD, &trade.ID, &trade.Amount, &trade.StartPrice, &trade.StopPrice, &trade.Threshold, &trade.Limit, &trade.CreatedAt, &trade.StartedAt, &trade.TryToStopAt, &trade.StoppedAt) {
			break
		}
		trades = append(trades, &trade)
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}

	iter = database.Session.Query("SELECT user_id, state, stock_id, srd, id, amount, start_price, stop_price, threshold, lim, created_at, started_at, try_to_stop_at, stopped_at FROM jeu_de_bourse.trades WHERE user_id = ? AND state = ? AND srd = ? ALLOW FILTERING;", userID, TradeStateTryToStop, srd).Iter()
	for {
		var trade Trade
		if !iter.Scan(&trade.UserID, &trade.State, &trade.StockID, &trade.SRD, &trade.ID, &trade.Amount, &trade.StartPrice, &trade.StopPrice, &trade.Threshold, &trade.Limit, &trade.CreatedAt, &trade.StartedAt, &trade.TryToStopAt, &trade.StoppedAt) {
			break
		}
		trades = append(trades, &trade)
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}

	return trades, nil
}

func GetOnGoingTradesByUserID(userID string) ([]*Trade, error) {
	var trades []*Trade

	// using 3 iters for each state
	iter := database.Session.Query("SELECT user_id, state, stock_id, srd, id, amount, start_price, stop_price, threshold, lim, created_at, started_at, try_to_stop_at, stopped_at FROM jeu_de_bourse.trades WHERE user_id = ? AND state = ? ALLOW FILTERING;", userID, TradeStateCreated).Iter()
	for {
		var trade Trade
		if !iter.Scan(&trade.UserID, &trade.State, &trade.StockID, &trade.SRD, &trade.ID, &trade.Amount, &trade.StartPrice, &trade.StopPrice, &trade.Threshold, &trade.Limit, &trade.CreatedAt, &trade.StartedAt, &trade.TryToStopAt, &trade.StoppedAt) {
			break
		}
		trades = append(trades, &trade)
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}

	iter = database.Session.Query("SELECT user_id, state, stock_id, srd, id, amount, start_price, stop_price, threshold, lim, created_at, started_at, try_to_stop_at, stopped_at FROM jeu_de_bourse.trades WHERE user_id = ? AND state = ? ALLOW FILTERING;", userID, TradeStateStarted).Iter()
	for {
		var trade Trade
		if !iter.Scan(&trade.UserID, &trade.State, &trade.StockID, &trade.SRD, &trade.ID, &trade.Amount, &trade.StartPrice, &trade.StopPrice, &trade.Threshold, &trade.Limit, &trade.CreatedAt, &trade.StartedAt, &trade.TryToStopAt, &trade.StoppedAt) {
			break
		}
		trades = append(trades, &trade)
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}

	iter = database.Session.Query("SELECT user_id, state, stock_id, srd, id, amount, start_price, stop_price, threshold, lim, created_at, started_at, try_to_stop_at, stopped_at FROM jeu_de_bourse.trades WHERE user_id = ? AND state = ? ALLOW FILTERING;", userID, TradeStateTryToStop).Iter()
	for {
		var trade Trade
		if !iter.Scan(&trade.UserID, &trade.State, &trade.StockID, &trade.SRD, &trade.ID, &trade.Amount, &trade.StartPrice, &trade.StopPrice, &trade.Threshold, &trade.Limit, &trade.CreatedAt, &trade.StartedAt, &trade.TryToStopAt, &trade.StoppedAt) {
			break
		}
		trades = append(trades, &trade)
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}

	return trades, nil
}
