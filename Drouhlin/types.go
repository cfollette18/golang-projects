package main

import (
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

type LoginResponse struct {
	Number int64  `json:"number"`
	Token  string `json:"token"`
}
type LoginRequest struct {
	Number   int64  `json:"number"`
	Password string `json:"password"`
}

type TransferRequest struct {
	ToAccount int     `json:"toAccount"`
	Amount    float64 `json:"amount"`
}

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
}

type Account struct {
	ID                int
	FirstName         string
	LastName          string
	Number            int64
	EncryptedPassword string `json:"-"`
	Balance           float64
	CreatedAt         time.Time
}

func (a *Account) ValidatePassword(pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(a.EncryptedPassword), []byte(pw)) == nil
}

func NewAccount(firstName, lastName, password string) (*Account, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &Account{
		FirstName:         firstName,
		LastName:          lastName,
		EncryptedPassword: string(encpw),
		Number:            int64(rand.Intn(10000000)),
		CreatedAt:         time.Now().UTC(),
	}, nil
}
