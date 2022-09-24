package usecase

import (
	"ca-boilerplate/domain"
	"ca-boilerplate/dto"
	"time"

	"github.com/google/uuid"
)

// Domain to DTO
func userDomainToDTOLogin(resp *dto.UserResponse, u *domain.User, r *domain.UserRole, token *string, refreshToken *string) {
	resp.UUID = u.UUID
	resp.Name = u.Name
	resp.Email = &u.Email
	resp.Role = r.Name
	resp.Token = token
	resp.RefreshToken = refreshToken
}

func userDomainToDTOFindUser(resp *dto.UserResponse, u *domain.User, r *domain.UserRole) {
	resp.UUID = u.UUID
	resp.Name = u.Name
	resp.Email = &u.Email
	respCreatedAt := time.UnixMicro(u.CreatedAt).UTC()
	resp.CreatedAt = &respCreatedAt

	if u.UpdatedAt != nil {
		respUpdatedAt := time.UnixMicro(*u.UpdatedAt).UTC()
		resp.UpdatedAt = &respUpdatedAt
	}

	if u.DeletedAt != nil {
		respDeletedAt := time.UnixMicro(*u.DeletedAt).UTC()
		resp.DeletedAt = &respDeletedAt
	}

	resp.Role = r.Name
}

func userDomainToDTORegister(resp *dto.UserResponse, u *domain.User, r *domain.UserRole) {
	resp.UUID = u.UUID
	resp.Name = u.Name
	resp.Email = &u.Email
	resp.Role = r.Name
	respCreatedAt := time.UnixMicro(u.CreatedAt).UTC()
	resp.CreatedAt = &respCreatedAt
}

// DTO to Domain
func userDTOtoDomainRegister(data *domain.User, req *dto.UserRegisterRequest) {
	data.UUID = uuid.New().String()
	data.Name = req.Name
	data.Email = req.Email
	data.Password = req.Password
	data.Phone = req.Phone
	data.Whatsapp = req.Whatsapp
	data.CreatedAt = time.Now().UnixMicro()
}
