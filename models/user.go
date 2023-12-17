package models

type User struct {
	ID           int
	Email        string
	Password     string
	SecretPhrase string
}
