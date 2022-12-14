package dbload

import (
	"log"

	"github.com/gocql/gocql"
)

func LoadUsers(session *gocql.Session) {
	// Generate tables if not exists
	if err := session.Query("CREATE TABLE IF NOT EXISTS jeu_de_bourse.users (id text PRIMARY KEY, username text, password text, email text, created_at timestamp, updated_at timestamp);").Exec(); err != nil {
		log.Println(err)
		return
	}

	// Creates indexes if not exists
	if err := session.Query("CREATE INDEX IF NOT EXISTS ON jeu_de_bourse.users (username);").Exec(); err != nil {
		log.Println(err)
		return
	}

	if err := session.Query("CREATE INDEX IF NOT EXISTS ON jeu_de_bourse.users (email);").Exec(); err != nil {
		log.Println(err)
		return
	}
}
