package controller

import (
	"net/http"

	"github.com/GAURAV/BookApiTask/database"
	"github.com/GAURAV/BookApiTask/models"
	"github.com/labstack/echo/v4"
)

func GetCategories(c echo.Context) error {
	var categories []models.Category
	if err := database.DB.Preload("Books").Find(&categories).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to fetch categories with books",
		})
	}
	return c.JSON(http.StatusOK, categories)
}

func CreateCategory(c echo.Context) error {
	category := new(models.Category)
	if err := c.Bind(category); err != nil {
		return err
	}
	database.DB.Create(&category)
	return c.JSON(http.StatusCreated, category)
}
