package controllers

import (
	"net/http"
	"praktikum/configs"
	books "praktikum/models/Books"
	"praktikum/resources"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetBooksController(c echo.Context) error {

	var GetBooks []books.Books
	var CountBooks int64

	BookException := configs.DB.Find(&GetBooks)
	BookCount := BookException.Count(&CountBooks)

	if BookException.Error != nil {
		return c.JSON(http.StatusInternalServerError, BookException.Error)
	}

	if BookCount.Error != nil {
		return c.JSON(http.StatusInternalServerError, BookCount.Error)
	}

	if CountBooks == 0 {
		return c.JSON(http.StatusNoContent, resources.BaseResponse{
			Status:  http.StatusNoContent,
			Message: "Book is empty",
			Data:    nil,
		})
	}

	BooksResponses := resources.MappingBooksResponse(GetBooks)

	return c.JSON(http.StatusOK, resources.BaseResponse{
		Status:  http.StatusOK,
		Message: "Success get all book",
		Data:    BooksResponses,
	})

}

func GetBookByIdController(c echo.Context) error {

	getId, IdException := strconv.Atoi(c.Param("id"))

	var (
		GetBook books.Books
	)

	if IdException != nil {
		return c.JSON(http.StatusBadRequest, resources.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: "Id Is Not Number",
			Data:    nil,
		})
	}

	BookException := configs.DB.First(&GetBook, "id = ?", getId)

	if BookException.Error != nil {
		return c.JSON(http.StatusInternalServerError, resources.BaseResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error When Find Book",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, resources.BaseResponse{
		Status:  http.StatusOK,
		Message: "Book Detail",
		Data:    GetBook,
	})
}

func CreateBookController(c echo.Context) error {

	var (
		GetRequestData books.BooksRequest
		BooksModel     books.Books
	)

	BindingException := c.Bind(&GetRequestData)

	if BindingException != nil {
		return c.JSON(http.StatusInternalServerError, BindingException.Error)
	}

	BooksModel = BooksModel.GetBooksData(GetRequestData)

	CreateException := configs.DB.Create(&BooksModel)

	if CreateException.Error != nil {
		return c.JSON(http.StatusInternalServerError, resources.BaseResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error When Create Book",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusCreated, resources.BaseResponse{
		Status:  http.StatusCreated,
		Message: "Book Created Success",
		Data:    BooksModel,
	})
}

func DeleteBookController(c echo.Context) error {

	getId, IdException := strconv.Atoi(c.Param("id"))

	if IdException != nil {
		return c.JSON(http.StatusBadRequest, resources.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: "Id Is Not Number",
			Data:    nil,
		})
	}

	Deleted := configs.DB.Delete(&books.Books{}, getId)

	if Deleted.Error != nil {
		return c.JSON(http.StatusInternalServerError, resources.BaseResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error when delete book",
			Data:    nil,
		})
	}

	if Deleted.RowsAffected == 0 {
		return c.JSON(http.StatusNoContent, resources.BaseResponse{
			Status:  http.StatusNoContent,
			Message: "Book not found",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, resources.BaseResponse{
		Status:  http.StatusNoContent,
		Message: "Book Deleted!",
		Data:    nil,
	})

}

func UpdateBookController(c echo.Context) error {

	getId, IdException := strconv.Atoi(c.Param("id"))

	var (
		GetRequestData books.BooksRequest
		BooksModel     books.Books
	)

	RequestBindingException := c.Bind(&GetRequestData)

	if IdException != nil {
		return c.JSON(http.StatusBadRequest, resources.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: "Id Is Not Number",
			Data:    nil,
		})
	}

	if RequestBindingException != nil {
		return c.JSON(http.StatusInternalServerError, RequestBindingException.Error)
	}

	BooksModel = BooksModel.GetBooksData(GetRequestData)

	updatedException := configs.DB.Where("id = ?", getId).Updates(&BooksModel)

	if updatedException.Error != nil {
		return c.JSON(http.StatusInternalServerError, updatedException.Error)
	}

	return c.JSON(http.StatusCreated, resources.BaseResponse{
		Status:  http.StatusCreated,
		Message: "Book update success",
		Data:    BooksModel,
	})

}
