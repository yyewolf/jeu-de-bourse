package dbload

import (
	"log"

	"github.com/gocql/gocql"
)

func LoadStock(session *gocql.Session) {
	// Generate tables if not exists
	if err := session.Query("CREATE TABLE IF NOT EXISTS jeu_de_bourse.stocks (id text PRIMARY KEY, symbol text, name text, exchange text, market text, srd boolean, loan boolean);").Exec(); err != nil {
		log.Println(err)
		return
	}

	// Generate history table if not exists
	if err := session.Query("CREATE TABLE IF NOT EXISTS jeu_de_bourse.stocks_history (trade_id text, stock_id text, price int, time timestamp, trade_type text, volume double, PRIMARY KEY ((trade_id), time)) WITH CLUSTERING ORDER BY (time DESC);").Exec(); err != nil {
		log.Println(err)
		return
	}

	// Creates indexes if not exists
	if err := session.Query("CREATE INDEX IF NOT EXISTS ON jeu_de_bourse.stocks (symbol);").Exec(); err != nil {
		log.Println(err)
		return
	}

	if err := session.Query("CREATE INDEX IF NOT EXISTS ON jeu_de_bourse.stocks (name);").Exec(); err != nil {
		log.Println(err)
		return
	}

	if err := session.Query("CREATE INDEX IF NOT EXISTS ON jeu_de_bourse.stocks (exchange);").Exec(); err != nil {
		log.Println(err)
		return
	}

	if err := session.Query("CREATE INDEX IF NOT EXISTS ON jeu_de_bourse.stocks (market);").Exec(); err != nil {
		log.Println(err)
		return
	}

	if err := session.Query("CREATE INDEX IF NOT EXISTS ON jeu_de_bourse.stocks (srd);").Exec(); err != nil {
		log.Println(err)
		return
	}

	if err := session.Query("CREATE INDEX IF NOT EXISTS ON jeu_de_bourse.stocks (loan);").Exec(); err != nil {
		log.Println(err)
		return
	}

	if err := session.Query("CREATE INDEX IF NOT EXISTS ON jeu_de_bourse.stocks_history (stock_id);").Exec(); err != nil {
		log.Println(err)
		return
	}

	if err := session.Query("CREATE INDEX IF NOT EXISTS ON jeu_de_bourse.stocks_history (time);").Exec(); err != nil {
		log.Println(err)
		return
	}

	if err := session.Query("CREATE INDEX IF NOT EXISTS ON jeu_de_bourse.stocks_history (trade_type);").Exec(); err != nil {
		log.Println(err)
		return
	}

	if err := session.Query("CREATE INDEX IF NOT EXISTS ON jeu_de_bourse.stocks_history (volume);").Exec(); err != nil {
		log.Println(err)
		return
	}

	if err := session.Query("CREATE INDEX IF NOT EXISTS ON jeu_de_bourse.stocks_history (price);").Exec(); err != nil {
		log.Println(err)
		return
	}
}
