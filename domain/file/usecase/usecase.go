package usecase

import (
	"ca-boilerplate/domain"
	"time"
)

type FileUploadUsecase struct {
	fileRepo       domain.FileRepositoryContract
	contextTimeout time.Duration
}

func NewFileUploadUsecase(repo domain.FileRepositoryContract, timeout time.Duration) domain.FileUsecaseContract {
	return &FileUploadUsecase{
		fileRepo:       repo,
		contextTimeout: timeout,
	}
}
