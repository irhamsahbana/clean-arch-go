package usecase

import (
	"ca-boilerplate/domain"
	"context"
)

func (u *bookUsecase) FindBooks(ctx context.Context) ([]domain.Book, int, error) {
	return u.bookRepository.FindBooks(ctx)
}
