package slice

import (
	"ca-boilerplate/domain"
	"context"
	"errors"
)

func (r *bookSliceRepository) CreateBook(ctx context.Context, book *domain.Book) (*domain.Book, int, error) {
	return book, 200, nil
}

func (r *bookSliceRepository) UpdateBook(ctx context.Context, id string, data *domain.Book) (*domain.Book, int, error) {
	for _, b := range books {
		if b.UUID == id {
			first := b

			first.Author = data.Author
			first.Description = data.Description
			first.Title = data.Title
			first.Year = data.Year
			return &first, 200, nil
		}
	}

	return nil, 404, errors.New("book not found")
}

func (r *bookSliceRepository) DeleteBook(ctx context.Context, id string) (*domain.Book, int, error) {
	for _, book := range books {
		if book.UUID == id {
			return &book, 200, nil
		}
	}

	return nil, 404, errors.New("book not found")
}
