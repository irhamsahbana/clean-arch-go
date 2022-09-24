package http

import (
	"ca-boilerplate/domain"
	"ca-boilerplate/dto"
	"ca-boilerplate/lib/http_response"
	"ca-boilerplate/lib/middleware"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type bookHandler struct {
	bookUsecase domain.BookUsecaseContract
}

func NewBookHandler(router *gin.Engine, usecase domain.BookUsecaseContract) {
	handler := &bookHandler{
		bookUsecase: usecase,
	}

	auth := router.Group("/", middleware.Auth)
	auth.POST("books", handler.CreateBook)
	auth.PUT("books/:id", handler.UpdateBook)
	auth.DELETE("books/:id", handler.DeleteBook)

	router.GET("books", handler.FindBooks)
	router.GET("books/:id", handler.FindBook)

}

func (h *bookHandler) FindBooks(c *gin.Context) {
	books, statusCode, err := h.bookUsecase.FindBooks(c)
	if err != nil {
		http_response.ReturnResponse(c, statusCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, statusCode, "success", books)
}

func (h *bookHandler) FindBook(c *gin.Context) {
	id := c.Param("id")

	book, statusCode, err := h.bookUsecase.FindBook(c, id)
	if err != nil {
		http_response.ReturnResponse(c, statusCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, statusCode, "success", book)
}

func (h *bookHandler) CreateBook(c *gin.Context) {
	var req dto.BookUpsertRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		http_response.ReturnResponse(c, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}

	ctx := context.Background()
	book, statusCode, err := h.bookUsecase.CreateBook(ctx, &req)
	if err != nil {
		http_response.ReturnResponse(c, statusCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, statusCode, "success", book)
}

func (h *bookHandler) UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var req dto.BookUpsertRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		http_response.ReturnResponse(c, 400, err.Error(), nil)
		return
	}

	ctx := context.Background()
	book, statusCode, err := h.bookUsecase.UpdateBook(ctx, id, &req)
	if err != nil {
		http_response.ReturnResponse(c, statusCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, statusCode, "success", book)
}

func (h *bookHandler) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	ctx := context.Background()
	result, statusCode, err := h.bookUsecase.DeleteBook(ctx, id)
	if err != nil {
		http_response.ReturnResponse(c, statusCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, statusCode, "success", result)
}
