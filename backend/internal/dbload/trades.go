package dbload

import (
	"log"

	"github.com/gocql/gocql"
)

func LoadTrades(session *gocql.Session) {
	// Generate tables if not exists
	if err := session.Query("CREATE TABLE IF NOT EXISTS jeu_de_bourse.trades (user_id text, state text, stock_id text, srd boolean, id text, amount int, start_price int, threshold int, lim int, stop_price int, created_at timestamp, started_at timestamp, try_to_stop_at timestamp, stopped_at timestamp, PRIMARY KEY ((user_id), stopped_at, state, stock_id, srd, id)) WITH CLUSTERING ORDER BY (stopped_at DESC, state DESC, stock_id DESC, srd DESC, id DESC);").Exec(); err != nil {
		log.Println(err)
	}

	// DROP ALL INDEXES
	if err := session.Query("DROP INDEX jeu_de_bourse.trades_state_idx;").Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query("DROP INDEX jeu_de_bourse.trades_stock_id_idx;").Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query("DROP INDEX jeu_de_bourse.trades_srd_idx;").Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query("DROP INDEX jeu_de_bourse.trades_id_idx;").Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query("DROP INDEX jeu_de_bourse.trades_amount_idx;").Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query("DROP INDEX jeu_de_bourse.trades_start_price_idx;").Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query("DROP INDEX jeu_de_bourse.trades_threshold_idx;").Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query("DROP INDEX jeu_de_bourse.trades_lim_idx;").Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query("DROP INDEX jeu_de_bourse.trades_stop_price_idx;").Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query("DROP INDEX jeu_de_bourse.trades_created_at_idx;").Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query("DROP INDEX jeu_de_bourse.trades_started_at_idx;").Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query("DROP INDEX jeu_de_bourse.trades_try_to_stop_at_idx;").Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query("DROP INDEX jeu_de_bourse.trades_stopped_at_idx;").Exec(); err != nil {
		log.Println(err)
	}

	// CREATE ALL INDEXES
	if err := session.Query("CREATE INDEX IF NOT EXISTS trades_state_idx ON jeu_de_bourse.trades (state);").Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query("CREATE INDEX IF NOT EXISTS trades_stock_id_idx ON jeu_de_bourse.trades (stock_id);").Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query("CREATE INDEX IF NOT EXISTS trades_srd_idx ON jeu_de_bourse.trades (srd);").Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query("CREATE INDEX IF NOT EXISTS trades_id_idx ON jeu_de_bourse.trades (id);").Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query("CREATE INDEX IF NOT EXISTS trades_amount_idx ON jeu_de_bourse.trades (amount);").Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query("CREATE INDEX IF NOT EXISTS trades_start_price_idx ON jeu_de_bourse.trades (start_price);").Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query("CREATE INDEX IF NOT EXISTS trades_threshold_idx ON jeu_de_bourse.trades (threshold);").Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query("CREATE INDEX IF NOT EXISTS trades_lim_idx ON jeu_de_bourse.trades (lim);").Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query("CREATE INDEX IF NOT EXISTS trades_stop_price_idx ON jeu_de_bourse.trades (stop_price);").Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query("CREATE INDEX IF NOT EXISTS trades_created_at_idx ON jeu_de_bourse.trades (created_at);").Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query("CREATE INDEX IF NOT EXISTS trades_started_at_idx ON jeu_de_bourse.trades (started_at);").Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query("CREATE INDEX IF NOT EXISTS trades_try_to_stop_at_idx ON jeu_de_bourse.trades (try_to_stop_at);").Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query("CREATE INDEX IF NOT EXISTS trades_stopped_at_idx ON jeu_de_bourse.trades (stopped_at);").Exec(); err != nil {
		log.Println(err)
	}
}
