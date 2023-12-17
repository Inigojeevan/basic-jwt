package repository

import (
	"errors"

	"github.com/Inigojeevan/jwt/models"
)

func FindCredentials(email string, password string) (*models.User, error) {
	if email == "sampleEmail@gmail.com" && password == "test1234" { //static
		return &models.User{
			ID:           1,
			Email:        "sampleEmail@gmail.com",
			Password:     "test1234",
			SecretPhrase: "Vannakam",
		}, nil
	}
	return nil, errors.New("User not found")
}
