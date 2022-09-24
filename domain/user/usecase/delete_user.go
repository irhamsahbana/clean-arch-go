package usecase

import (
	"ca-boilerplate/dto"
	"context"
)

func (u *userUsecase) DeleteUser(c context.Context, id string) (*dto.UserResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	// delete user
	result, code, err := u.userRepo.DeleteUser(ctx, id)
	if err != nil {
		return nil, code, err
	}

	// get role
	role, code, err := u.userRoleRepo.FindUserRoleBy(ctx, "uuid", result.RoleUUID, false)
	if err != nil {
		return nil, code, err
	}

	var resp dto.UserResponse
	userDomainToDTOFindUser(&resp, result, role)

	return &resp, code, nil
}
