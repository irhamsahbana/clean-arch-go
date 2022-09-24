package domain

import (
	"ca-boilerplate/dto"
	"context"
)

type Book struct {
	UUID        string `json:"uuid"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Year        int    `json:"year"`
}

type BookUsecaseContract interface {
	FindBooks(ctx context.Context) ([]Book, int, error)
	FindBook(ctx context.Context, id string) (*Book, int, error)
	CreateBook(ctx context.Context, book *dto.BookUpsertRequest) (*Book, int, error)
	UpdateBook(ctx context.Context, id string, book *dto.BookUpsertRequest) (*Book, int, error)
	DeleteBook(ctx context.Context, id string) (*Book, int, error)
}

type BookRepositoryContract interface {
	FindBooks(ctx context.Context) ([]Book, int, error)
	FindBookBy(ctx context.Context, field, value string) (*Book, int, error)
	CreateBook(ctx context.Context, book *Book) (*Book, int, error)
	UpdateBook(ctx context.Context, id string, book *Book) (*Book, int, error)
	DeleteBook(ctx context.Context, id string) (*Book, int, error)
}
