package handlers

import (
	"time"

	"github.com/Inigojeevan/jwt/config"
	"github.com/Inigojeevan/jwt/models"
	"github.com/Inigojeevan/jwt/repository"
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
)

func Login(c *fiber.Ctx) error {
	loginRequest := new(models.LoginRequest)

	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	user, err := repository.FindCredentials(loginRequest.Email, loginRequest.Password)

	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	day := time.Hour * 24

	claims := jtoken.MapClaims{
		"ID":           user.ID,
		"Email":        user.Email,
		"SecretPhrase": user.SecretPhrase,
		"Expiry_At":    time.Now().Add(day).Unix(),
	}

	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims) //creates the token

	t, err := token.SignedString([]byte(config.Secret))

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.JSON(fiber.Map{"Token": t})
}

func ProtectedRoute(c *fiber.Ctx) error {
	user := c.Locals("user").(*jtoken.Token)
	claims := user.Claims.(jtoken.MapClaims)
	email := claims["Email"].(string)
	SecretPhrase := claims["SecretPhrase"].(string)

	return c.SendString("welcome " + email + " " + SecretPhrase)
}
