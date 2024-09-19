package handler

import (
	"fmt"
	"github.com/Ahmad-mufied/broker-service_gc2p3/config"
	"github.com/Ahmad-mufied/broker-service_gc2p3/constans"
	"github.com/Ahmad-mufied/broker-service_gc2p3/model"
	"github.com/Ahmad-mufied/broker-service_gc2p3/pb"
	"github.com/Ahmad-mufied/broker-service_gc2p3/utils"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type BookHandler struct {
	bookService pb.BookServiceClient
}

func NewBookHandler(client pb.BookServiceClient) *BookHandler {
	return &BookHandler{
		bookService: client,
	}
}

func (h *BookHandler) CreateBook(c echo.Context) error {
	book := new(model.Book)

	err := c.Bind(&book)
	if err != nil {
		return utils.HandleError(c, constans.ErrBadRequest, "Invalid request body")
	}

	// Validate the order struct
	err = config.Validator.Struct(book)
	if err != nil {
		// Format the validation errors
		errors := utils.FormatValidationErrors(err)
		return utils.HandleValidationError(c, errors)
	}

	req := pb.CreateBookRequest{
		Title:       book.Title,
		Author:      book.Author,
		PublishDate: book.PublishData,
		Status:      book.Status,
	}

	res, err := h.bookService.Create(c.Request().Context(), &req)
	if err != nil {
		return utils.HandleError(c, constans.ErrInternalServerError, err.Error())
	}

	if !res.Success {
		return utils.HandleError(c, constans.ErrInternalServerError, "Failed to create book")
	}

	book.Id = res.Id

	return c.JSON(http.StatusCreated, model.JsonResponse{Status: "success", Message: "Book created", Data: book})
}

func (h *BookHandler) GetBookById(c echo.Context) error {
	bookId := c.Param("id")

	// check bookId is valid
	if bookId == "" {
		return utils.HandleError(c, constans.ErrBadRequest, "Please provide book id as parameter")
	}

	log.Printf("Book ID: %s\n", bookId)

	req := pb.GetBookByIdRequest{
		Id: bookId,
	}

	res, err := h.bookService.GetBookById(c.Request().Context(), &req)
	if err != nil {
		return utils.HandleError(c, constans.ErrInternalServerError, fmt.Sprintf("Failed to get book by id: %s Eroor: %s", bookId, err.Error()))
	}

	log.Printf("Book exist: %v\n", res.Success)

	if res.Success == false {
		return utils.HandleError(c, constans.ErrNotFound, fmt.Sprintf("Book with id: %s not found", bookId))
	}

	book := model.Book{
		Id:          bookId,
		Title:       res.Title,
		Author:      res.Author,
		PublishData: res.PublishDate,
		Status:      res.Status,
	}

	return c.JSON(http.StatusOK, model.JsonResponse{Status: "success", Message: "Book found", Data: book})
}

func (h *BookHandler) GetBookByTitle(c echo.Context) error {
	title := c.Param("title")

	req := pb.GetBookByTitleRequest{
		Title: title,
	}

	res, err := h.bookService.GetBookByTitle(c.Request().Context(), &req)
	if err != nil {
		return utils.HandleError(c, constans.ErrInternalServerError, err.Error())
	}

	book := model.Book{
		Id:          res.Id,
		Title:       title,
		Author:      res.Author,
		PublishData: res.PublishDate,
		Status:      res.Status,
	}

	return c.JSON(http.StatusOK, model.JsonResponse{Status: "success", Message: "Book found", Data: book})
}

func (h *BookHandler) UpdateBook(c echo.Context) error {

	bookId := c.Param("id")

	// check bookId is valid
	if bookId == "" {
		return utils.HandleError(c, constans.ErrBadRequest, "Please provide book id as parameter")
	}

	// Check if book exist
	check, err := h.bookService.Check(c.Request().Context(), &pb.CheckBookRequest{Id: bookId})
	if err != nil {
		return utils.HandleError(c, constans.ErrInternalServerError, "Failed to check book")
	}

	if !check.Exist {
		return utils.HandleError(c, constans.ErrNotFound, fmt.Sprintf("Book with id: %s not found", bookId))
	}

	book := new(model.Book)

	err = c.Bind(&book)
	if err != nil {
		return utils.HandleError(c, constans.ErrBadRequest, "Invalid request body")
	}

	// Validate the order struct
	err = config.Validator.Struct(book)
	if err != nil {
		// Format the validation errors
		errors := utils.FormatValidationErrors(err)
		return utils.HandleValidationError(c, errors)
	}

	req := pb.UpdateBookRequest{
		Id:          bookId,
		Title:       book.Title,
		Author:      book.Author,
		PublishDate: book.PublishData,
		Status:      book.Status,
	}

	res, err := h.bookService.Update(c.Request().Context(), &req)
	if err != nil {
		return utils.HandleError(c, constans.ErrInternalServerError, err.Error())
	}

	if !res.Success {
		return utils.HandleError(c, constans.ErrInternalServerError, "Failed to update book")
	}

	return c.JSON(http.StatusOK, model.JsonResponse{Status: "success", Message: "Book updated"})
}

func (h *BookHandler) DeleteBook(c echo.Context) error {
	bookId := c.Param("id")

	// check bookId is valid
	if bookId == "" {
		return utils.HandleError(c, constans.ErrBadRequest, "Please provide book id as parameter")
	}

	// Check if book exist
	check, err := h.bookService.Check(c.Request().Context(), &pb.CheckBookRequest{Id: bookId})
	if err != nil {
		return utils.HandleError(c, constans.ErrInternalServerError, "Failed to check book")
	}

	if !check.Exist {
		return utils.HandleError(c, constans.ErrNotFound, fmt.Sprintf("Book with id: %s not found", bookId))
	}

	req := pb.DeleteBookRequest{
		Id: bookId,
	}

	// Check if book exist
	isExist, err := h.bookService.GetBookById(c.Request().Context(), &pb.GetBookByIdRequest{Id: bookId})
	if err != nil {
		return utils.HandleError(c, constans.ErrNotFound, fmt.Sprintf("Book with id: %s not found", bookId))
	}

	if isExist.Success == false {
		return utils.HandleError(c, constans.ErrNotFound, fmt.Sprintf("Book with id: %s not found", bookId))
	}

	res, err := h.bookService.Delete(c.Request().Context(), &req)
	if err != nil {
		return utils.HandleError(c, constans.ErrInternalServerError, err.Error())
	}

	if !res.Success {
		return utils.HandleError(c, constans.ErrInternalServerError, "Failed to delete book")
	}

	return c.JSON(http.StatusOK, model.JsonResponse{Status: "success", Message: "Book deleted"})
}
