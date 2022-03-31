package book

import (
	"be7/layered/delivery/helper"
	_entities "be7/layered/entities"
	_bookUseCase "be7/layered/usecase/book"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookHandler struct {
	bookUseCase _bookUseCase.BookUseCaseInterface
}

func NewBookHandler(bookUseCase _bookUseCase.BookUseCaseInterface) *BookHandler {
	return &BookHandler{
		bookUseCase: bookUseCase,
	}
}

func (bh *BookHandler) GetAllHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		books, err := bh.bookUseCase.GetAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed fetch data"))
		}
		if len(books) == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Data not exist"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get all data", books))
	}
}

func (bh *BookHandler) GetBookHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idn := c.Param("id")
		id, _ := strconv.Atoi(idn)
		books, rows, err := bh.bookUseCase.GetBook(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not exist"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("succes get data", books))
	}
}

func (bh *BookHandler) DeleteBookHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idn := c.Param("id")
		id, _ := strconv.Atoi(idn)
		_, rows, err := bh.bookUseCase.GetBook(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not exist"))
		}
		_, err = bh.bookUseCase.DeleteBook(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed delete data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success delete data"))
	}
}
func (bh *BookHandler) CreateBookHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var book _entities.Book
		c.Bind(&book)
		books, err := bh.bookUseCase.CreateBook(book)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed create data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success create data", books))
	}
}

func (bh *BookHandler) UpdatedBookHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idn := c.Param("id")
		id, _ := strconv.Atoi(idn)
		books, rows, err := bh.bookUseCase.GetBook(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not exist"))
		}
		c.Bind(&books)
		books, err = bh.bookUseCase.UpdatedBook(books, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed edit data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success edit data", books))
	}
}
