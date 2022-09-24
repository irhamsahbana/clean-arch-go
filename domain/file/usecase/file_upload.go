package usecase

import (
	"ca-boilerplate/domain"
	"ca-boilerplate/dto"
	custom "ca-boilerplate/lib/custom_type"
	"ca-boilerplate/lib/validator"
	"context"
	"errors"
	"mime/multipart"
	"net/http"
	"path/filepath"
)

func (u *FileUploadUsecase) UploadFile(c context.Context, file *multipart.FileHeader, req *dto.UploadFileRequest) (*dto.UploadFileResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	if err := validator.IsUUID(req.FileableUUID); err != nil {
		return nil, http.StatusUnprocessableEntity, err
	}

	if err := validateFileableType(req.FileableType); err != nil {
		return nil, http.StatusUnprocessableEntity, err
	}

	var data domain.File
	data.UUID = req.UUID
	data.FileableUUID = req.FileableUUID
	data.FileableType = req.FileableType
	data.Ext = filepath.Ext(file.Filename)
	data.Path = "storage/" + string(data.FileableType) + "/" + data.UUID + data.Ext

	result, code, err := u.fileRepo.UploadFile(ctx, &data)
	if err != nil {
		return nil, code, err
	}

	var resp dto.UploadFileResponse
	resp.UUID = result.UUID
	resp.FileableUUID = result.FileableUUID
	resp.FileableType = result.FileableType
	resp.Ext = result.Ext
	resp.Path = result.Path

	// save file *multipart.FileHeader to storage

	return &resp, http.StatusOK, nil
}

func validateFileableType(ft custom.FileableType) error {
	switch ft {
	case "item_categories.items":
		return nil
	default:
		return errors.New("invalid fileable type")
	}
}
