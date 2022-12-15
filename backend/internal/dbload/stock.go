package dbload

import (
	"log"

	"github.com/gocql/gocql"
)

func LoadStock(session *gocql.Session) {
	// Generate tables if not exists
	if err := session.Query("CREATE TABLE IF NOT EXISTS jeu_de_bourse.stocks (id text PRIMARY KEY, symbol text, name text, exchange text, market text, srd boolean, loan boolean);").Exec(); err != nil {
		log.Println(err)
	}

	// Generate history table if not exists
	if err := session.Query("CREATE TABLE IF NOT EXISTS jeu_de_bourse.stocks_history (trade_id text, stock_id text, price int, time timestamp, trade_type text, volume double, PRIMARY KEY ((stock_id), time, trade_id)) WITH CLUSTERING ORDER BY (time DESC);").Exec(); err != nil {
		log.Println(err)
	}

	// DROP ALL INDEXES
	if err := session.Query("DROP INDEX jeu_de_bourse.stocks_symbol_idx;").Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query("DROP INDEX jeu_de_bourse.stocks_name_idx;").Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query("DROP INDEX jeu_de_bourse.stocks_exchange_idx;").Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query("DROP INDEX jeu_de_bourse.stocks_market_idx;").Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query("DROP INDEX jeu_de_bourse.stocks_srd_idx;").Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query("DROP INDEX jeu_de_bourse.stocks_loan_idx;").Exec(); err != nil {
		log.Println(err)
	}

	// if err := session.Query("DROP INDEX jeu_de_bourse.stocks_history_time_idx;").Exec(); err != nil {
	// 	log.Println(err)
	// }

	if err := session.Query("DROP INDEX jeu_de_bourse.stocks_history_trade_type_idx;").Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query("DROP INDEX jeu_de_bourse.stocks_history_volume_idx;").Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query("DROP INDEX jeu_de_bourse.stocks_history_price_idx;").Exec(); err != nil {
		log.Println(err)
	}

	// Creates indexes if not exists
	if err := session.Query("CREATE CUSTOM INDEX IF NOT EXISTS ON jeu_de_bourse.stocks (symbol) USING 'org.apache.cassandra.index.sasi.SASIIndex' WITH OPTIONS = { 'mode': 'CONTAINS', 'case_sensitive': 'false' };").Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query("CREATE CUSTOM INDEX IF NOT EXISTS ON jeu_de_bourse.stocks (name) USING 'org.apache.cassandra.index.sasi.SASIIndex' WITH OPTIONS = { 'mode': 'CONTAINS', 'case_sensitive': 'false' };").Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query("CREATE INDEX IF NOT EXISTS ON jeu_de_bourse.stocks (exchange);").Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query("CREATE INDEX IF NOT EXISTS ON jeu_de_bourse.stocks (market);").Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query("CREATE INDEX IF NOT EXISTS ON jeu_de_bourse.stocks (srd);").Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query("CREATE INDEX IF NOT EXISTS ON jeu_de_bourse.stocks (loan);").Exec(); err != nil {
		log.Println(err)
	}

	// if err := session.Query("CREATE INDEX IF NOT EXISTS ON jeu_de_bourse.stocks_history (time);").Exec(); err != nil {
	// 	log.Println(err)
	// }

	if err := session.Query("CREATE INDEX IF NOT EXISTS ON jeu_de_bourse.stocks_history (trade_type);").Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query("CREATE INDEX IF NOT EXISTS ON jeu_de_bourse.stocks_history (volume);").Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query("CREATE INDEX IF NOT EXISTS ON jeu_de_bourse.stocks_history (price);").Exec(); err != nil {
		log.Println(err)
	}
}
