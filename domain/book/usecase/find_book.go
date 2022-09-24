package usecase

import (
	"ca-boilerplate/domain"
	"context"
)

func (u *bookUsecase) FindBook(ctx context.Context, id string) (*domain.Book, int, error) {
	return u.bookRepository.FindBookBy(ctx, "uuid", id)
}
