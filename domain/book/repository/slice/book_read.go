package slice

import (
	"ca-boilerplate/domain"
	"context"
	"errors"
)

func (r *bookSliceRepository) FindBookBy(ctx context.Context, field, value string) (*domain.Book, int, error) {

	for _, book := range books {
		if book.UUID == value {
			return &book, 200, nil
		}
	}

	return nil, 404, errors.New("book not found")
}

func (r *bookSliceRepository) FindBooks(ctx context.Context) ([]domain.Book, int, error) {
	return books, 200, nil
}
