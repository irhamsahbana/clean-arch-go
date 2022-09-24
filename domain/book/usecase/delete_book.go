package usecase

import (
	"ca-boilerplate/domain"
	"context"
)

func (u *bookUsecase) DeleteBook(ctx context.Context, id string) (*domain.Book, int, error) {
	return u.bookRepository.DeleteBook(ctx, id)
}
