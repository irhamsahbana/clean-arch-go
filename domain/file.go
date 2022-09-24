package domain

import (
	"ca-boilerplate/dto"
	custom "ca-boilerplate/lib/custom_type"
	"context"
	"mime/multipart"
)

type File struct {
	UUID         string              `bson:"uuid"`
	FileableUUID string              `bson:"fileable_uuid"`
	FileableType custom.FileableType `bson:"fileable_type"`
	Ext          string              `bson:"ext"`
	Path         string              `bson:"path"`
	CreatedAt    int64               `bson:"created_at"`
	UpdatedAt    *int64              `bson:"updated_at,omitempty"`
	DeletedAt    *int64              `bson:"deleted_at,omitempty"`
}

type FileUsecaseContract interface {
	UploadFile(ctx context.Context, file *multipart.FileHeader, req *dto.UploadFileRequest) (*dto.UploadFileResponse, int, error)
}

type FileRepositoryContract interface {
	UploadFile(ctx context.Context, data *File) (*File, int, error)
}
