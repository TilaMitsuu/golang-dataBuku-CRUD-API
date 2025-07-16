package controllers

import (
	"time"

	"book-api/config"
	"book-api/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("secretkey")

func Register(c *fiber.Ctx) error {
	type Request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var body Request
	if err := c.BodyParser(&body); err != nil {
		return err
	}

	// Hash password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(body.Password), 14)

	user := models.User{
		Username: body.Username,
		Password: string(hashedPassword),
	}

	result := config.DB.Create(&user)
	if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Username already exists",
		})
	}

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	type Request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var body Request
	if err := c.BodyParser(&body); err != nil {
		return err
	}

	var user models.User
	config.DB.Where("username = ?", body.Username).First(&user)

	if user.ID == 0 {
		return c.Status(400).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid password",
		})
	}

	// Buat token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"token": tokenString,
	})
}

func JwtSecret() []byte {
	return jwtSecret
}