package usecase

import (
	"ca-boilerplate/domain"
	"ca-boilerplate/dto"
	"context"
)

func (u *bookUsecase) UpdateBook(ctx context.Context, id string, req *dto.BookUpsertRequest) (*domain.Book, int, error) {
	var data domain.Book
	data.Title = req.Title
	data.Description = req.Description
	data.Author = req.Author
	data.Year = req.Year
	return u.bookRepository.UpdateBook(ctx, id, &data)
}
