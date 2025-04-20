package controller

import (
	"net/http"

	"github.com/GAURAV/BookApiTask/database"
	"github.com/GAURAV/BookApiTask/models"
	"github.com/labstack/echo/v4"
)

func GetBooks(c echo.Context) error {
	var books []models.Book
	database.DB.Preload("Category").Find(&books)
	return c.JSON(http.StatusOK, books)
}

func CreateBook(c echo.Context) error {
	var req models.Book

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid request body"})
	}

	var category models.Category
	if err := database.DB.First(&category, req.CategoryID).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid category ID"})
	}

	if err := database.DB.Create(&req).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to create book"})
	}

	return c.JSON(http.StatusCreated, req)
}

func UpdateBook(c echo.Context) error {
	id := c.Param("id")
	var book models.Book
	database.DB.First(&book, id)
	if book.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Book not found"})
	}
	if err := c.Bind(&book); err != nil {
		return err
	}
	database.DB.Save(&book)
	return c.JSON(http.StatusOK, book)
}

func DeleteBook(c echo.Context) error {
	id := c.Param("id")
	var book models.Book
	database.DB.Delete(&book, id)
	return c.JSON(http.StatusOK, echo.Map{"message": "Book deleted"})
}
