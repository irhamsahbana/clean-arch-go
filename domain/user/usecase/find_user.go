package usecase

import (
	"ca-boilerplate/dto"
	"context"
	"net/http"
)

func (u *userUsecase) FindUser(c context.Context, id string, withTrashed bool) (*dto.UserResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	result, code, err := u.userRepo.FindUserBy(ctx, "uuid", id, withTrashed)
	if err != nil {
		return nil, code, err
	}

	role, code, err := u.userRoleRepo.FindUserRoleBy(ctx, "uuid", result.RoleUUID, withTrashed)
	if err != nil {
		return nil, code, err
	}

	var resp dto.UserResponse
	userDomainToDTOFindUser(&resp, result, role)

	return &resp, http.StatusOK, nil
}
