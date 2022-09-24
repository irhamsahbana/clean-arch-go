package usecase

import (
	"ca-boilerplate/domain"
	"ca-boilerplate/dto"
	"context"
)

func (u *bookUsecase) CreateBook(ctx context.Context, req *dto.BookUpsertRequest) (*domain.Book, int, error) {
	var data domain.Book
	data.Title = req.Title
	data.Description = req.Description
	data.Author = req.Author
	data.Year = req.Year
	data.UUID = "123"

	return u.bookRepository.CreateBook(ctx, &data)
}
