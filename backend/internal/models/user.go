package models

import (
	"jeu-de-bourse/internal/database"
	"jeu-de-bourse/internal/env"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (u *User) CheckPasswordHash(attempt string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(attempt))
	return err == nil
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	err := database.Session.Query("SELECT id, username, password, email, created_at, updated_at FROM jeu_de_bourse.users WHERE username = ?", username).
		Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	return &user, err
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	err := database.Session.Query("SELECT id, username, password, email, created_at, updated_at FROM jeu_de_bourse.users WHERE email = ?", email).
		Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	return &user, err
}

func GetUserByID(id string) (*User, error) {
	var user User
	err := database.Session.Query("SELECT id, username, password, email, created_at, updated_at FROM jeu_de_bourse.users WHERE id = ?", id).
		Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	return &user, err
}

func (u *User) Create() error {
	err := database.Session.Query("INSERT INTO jeu_de_bourse.users (id, username, password, email, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)", u.ID, u.Username, u.Password, u.Email, u.CreatedAt, u.UpdatedAt).Exec()
	return err
}

func (u *User) DisponibleComptant() (amount int, err error) {
	// Start Comptant + Total des gains - Total des pertes - Investissement
	completedTrades, _ := GetCompletedTradesByUserID(u.ID, false)

	var total int
	for _, trade := range completedTrades {
		ratio := (1.0 - float64(trade.StartPrice)/float64(trade.StopPrice))
		total += int(float64(trade.Amount) * ratio)
	}

	amount = env.GlobalConfig.StartComptant + total

	ongoingTrades, _ := GetOnGoingTradesByUserIDWithSRD(u.ID, false)
	for _, trade := range ongoingTrades {
		total -= trade.Amount
	}
	amount += total
	return amount, nil
}

func (u *User) DisponibleSRD() (amount int, err error) {
	// Start SRD + Total des gains - Total des pertes - Investissement
	completedTrades, _ := GetCompletedTradesByUserID(u.ID, true)

	var total int
	for _, trade := range completedTrades {
		ratio := (1.0 - float64(trade.StartPrice)/float64(trade.StopPrice))
		total += int(float64(trade.Amount) * ratio)
	}

	amount = env.GlobalConfig.StartSRD + total

	ongoingTrades, _ := GetOnGoingTradesByUserIDWithSRD(u.ID, true)
	for _, trade := range ongoingTrades {
		total -= trade.Amount
	}
	amount += total
	return amount, nil
}

func (u *User) Investis() (amount int, err error) {
	ongoingTrades, err := GetOnGoingTradesByUserID(u.ID)
	for _, trade := range ongoingTrades {
		amount += trade.Amount
	}
	return amount, nil
}
func (u *User) Valorisation() (amount int, err error) {
	// Start Comptant + Total des gains - Total des pertes - Investissement
	completedTrades, _ := GetCompletedTradesByUserID(u.ID, false)

	var total int
	for _, trade := range completedTrades {
		ratio := (1.0 - float64(trade.StartPrice)/float64(trade.StopPrice))
		total += int(float64(trade.Amount) * ratio)
	}
	return env.GlobalConfig.StartComptant + total, nil
}
