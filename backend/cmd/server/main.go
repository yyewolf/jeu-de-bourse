package main

import (
	"jeu-de-bourse/api/v1/router"
	"jeu-de-bourse/internal/env"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var secret = []byte(env.GlobalConfig.CookieKey)

func main() {
	// Loads the default config
	app := gin.Default()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:8080"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin, X-Requested-With, Content-Type, Accept"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Setup the session
	store := cookie.NewStore(secret)
	store.Options(sessions.Options{
		MaxAge: 10 * 60 * 60 * 24,
	})
	app.Use(sessions.Sessions(env.GlobalConfig.CookieName, store))

	// Setup the routes
	router.Route(app)

	app.Run(":8080")
}
