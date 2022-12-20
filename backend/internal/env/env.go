package env

import (
	"log"
	"os"
	"strconv"

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

	// Settings
	StartComptant int
	StartSRD      int
	Fee           float64
}

func Getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func Getenvint(key string, def int) int {
	if v := os.Getenv(key); v != "" {
		d, err := strconv.Atoi(v)
		if err != nil {
			return def
		}
		return d
	}
	return def
}

func Getenvfloat(key string, def float64) float64 {
	if v := os.Getenv(key); v != "" {
		d, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return def
		}
		return d
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
		GlobalConfig.StartComptant = Getenvint("JDB_START_COMPTANT", 300000000)
		GlobalConfig.StartSRD = Getenvint("JDB_START_SRD", 700000000)
		GlobalConfig.Fee = Getenvfloat("JDB_FEE", 0.001)
	}
}
