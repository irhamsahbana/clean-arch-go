package dto

import (
	custom "ca-boilerplate/lib/custom_type"
)

type UploadFileRequest struct {
	UUID         string              `json:"uuid"`
	FileableUUID string              `json:"fileable_uuid"`
	FileableType custom.FileableType `json:"fileable_type"`
}

type UploadFileResponse struct {
	UUID         string              `json:"uuid"`
	FileableUUID string              `json:"fileable_uuid"`
	FileableType custom.FileableType `json:"fileable_type"`
	Url          string              `json:"url"`
	Path         string              `json:"path"`
	Ext          string              `json:"ext"`
}
