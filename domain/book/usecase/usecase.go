package usecase

import "ca-boilerplate/domain"

type bookUsecase struct {
	bookRepository domain.BookRepositoryContract
}

func NewBookUsecase(bookRepo domain.BookRepositoryContract) domain.BookUsecaseContract {
	return &bookUsecase{
		bookRepository: bookRepo,
	}
}
