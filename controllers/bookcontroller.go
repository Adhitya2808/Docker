package controllers

import (
	"net/http"
	"Praktikum/configs"
	"Praktikum/model"
	"strconv"
	"github.com/labstack/echo/v4"
)

func GetBookController(c echo.Context) error {
    var books []model.Book
    if err := configs.DB.Find(&books).Error; err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "success get all books",
        "books":   books,
    })
}

func GetBookidController(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid book ID")
    }
    var books []model.Book
    if err := configs.DB.First(&books, id).Error; err != nil {
        return echo.NewHTTPError(http.StatusNotFound, "Books not found")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "success get books",
        "books":    books,
    })
}

func CreateBookController(c echo.Context) error {
    books := model.Book{}
    c.Bind(&books)
    if err := configs.DB.Create(&books).Error; err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "success create new books",
        "books":    books,
    })
}

func DeleteBookController(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid book ID")
    }
	var books model.Book
    if err := configs.DB.Delete(&books, id).Error; err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete Books")
    }
    return c.JSON(http.StatusOK, map[string]string{
        "message": "Books deleted",
    })
}

func UpdateBookController(c echo.Context) error {
    var update_books model.Book
    if err := c.Bind(&update_books); err != nil{
        return c.JSON(http.StatusBadRequest, model.BaseResponse{
            Status: false,
            Message: "Books Not Found",
            Data : nil,
        })
    }
    bookID, err := strconv.Atoi(c.Param("id"))
    if err != nil{
        return c.JSON(http.StatusBadRequest, model.BaseResponse{
            Status: false,
            Message: "Invalid book ID",
            Data : nil,
        })
    }
    exsitingbook := model.Book{}
    err = configs.DB.First(&exsitingbook, bookID).Error
    if err != nil {
            return c.JSON(http.StatusInternalServerError, model.BaseResponse{
            Status: false,
            Message: "Failed to Fetch Book",
            Data : nil,
        })
    }
    exsitingbook.Judul = update_books.Judul
    exsitingbook.Penulis = update_books.Penulis
    exsitingbook.Penerbit = update_books.Penerbit

    err = configs.DB.Save(&exsitingbook).Error
    if err != nil{
        return c.JSON(http.StatusInternalServerError, model.BaseResponse{
            Status: false,
            Message: "Failed to update book",
            Data : nil,
        })
    }
    return c.JSON(http.StatusOK, model.BaseResponse{
            Status: true,
            Message: "Update Successfully",
            Data : nil,
    })
}