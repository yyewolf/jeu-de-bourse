package database

import (
	"jeu-de-bourse/internal/dbload"
	"jeu-de-bourse/internal/env"
	"log"
	"time"

	"github.com/gocql/gocql"
)

var Session *gocql.Session

func init() {
	// connect to the cluster
	cluster := gocql.NewCluster(env.GlobalConfig.CassandraHost)
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 3
	cluster.ConnectTimeout = time.Second * 10
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: env.GlobalConfig.CassandraUsername, Password: env.GlobalConfig.CassandraPassword}
	session, err := cluster.CreateSession()
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("[JDB-DB] Connected to cassandra.")
	Session = session

	// Generate keyspace if not exists
	if err := session.Query("CREATE KEYSPACE IF NOT EXISTS jeu_de_bourse WITH REPLICATION = {'class':'SimpleStrategy', 'replication_factor':1};").Exec(); err != nil {
		log.Println(err)
		return
	}

	dbload.LoadUsers(session)
	dbload.LoadStock(session)
	dbload.LoadTrades(session)

	log.Println("[JDB-DB] Created keyspace and tables.")
}
