package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	GlobalConfig Config
)

type Config struct {
	Mode string

	// Database
	CassandraHost     string
	CassandraUsername string
	CassandraPassword string

	// Recapcha
	RecaptchaSecret string

	// Cookie
	CookieName string
	CookieKey  string
}

func Getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func init() {
	MODE := Getenv("JDB_MODE", "dev")
	GlobalConfig.Mode = MODE
	if MODE == "dev" {
		godotenv.Load("development.env")
		log.Println("[JDB-ENV] Loaded development.env")
		GlobalConfig.CassandraHost = Getenv("JDB_CASSANDRA_HOST", "127.0.0.1")
		GlobalConfig.CassandraUsername = Getenv("JDB_CASSANDRA_USERNAME", "cassandra")
		GlobalConfig.CassandraPassword = Getenv("JDB_CASSANDRA_PASSWORD", "cassandra")
		GlobalConfig.RecaptchaSecret = Getenv("JDB_RECAPTCHA_SECRET", "6LcZ8uIUAAAAAER2Z8Z8Z8Z8Z8Z8Z8Z8Z8Z8Z8Z8")
		GlobalConfig.CookieName = Getenv("JDB_COOKIE_NAME", "jeu-de-bourse")
		GlobalConfig.CookieKey = Getenv("JDB_COOKIE_KEY", "jeu-de-bourse")
	}
}
