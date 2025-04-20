package controller

import (
	"net/http"
	"time"

	"github.com/GAURAV/BookApiTask/database"
	"github.com/GAURAV/BookApiTask/models"
	"github.com/GAURAV/BookApiTask/utils"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	hashedPassword, _ := utils.HashPassword(user.PassWord)
	user.PassWord = hashedPassword
	result := database.DB.Create(&user)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Registration failed"})
	}
	return c.JSON(http.StatusCreated, user)
}

func Login(c echo.Context) error {
	var req models.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid request body"})
	}

	var user models.User
	if err := database.DB.First(&user, "user_name = ?", req.Username).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Invalid credentials"})
	}

	if !utils.CheckPassword(user.PassWord, req.Password) {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Invalid credentials"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	t, _ := token.SignedString([]byte("secret"))

	return c.JSON(http.StatusOK, echo.Map{"token": t})
}
