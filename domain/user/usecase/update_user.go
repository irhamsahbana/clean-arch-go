package usecase

import (
	"ca-boilerplate/dto"
	passwordhandler "ca-boilerplate/lib/password_handler"
	"context"
	"net/http"
	"time"
)

func (u *userUsecase) UpdateUser(c context.Context, id string, req *dto.UserUpdateRequest) (*dto.UserResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	user, code, err := u.userRepo.FindUserBy(ctx, "uuid", id, false)
	if err != nil {
		return nil, code, err
	}

	user.Name = req.Name
	user.Email = req.Email
	user.Phone = req.Phone
	user.Whatsapp = req.Whatsapp
	if req.Password != nil {
		// hash password
		pass, err := passwordhandler.HashPassword(*req.Password)
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}

		user.Password = pass
	}
	updateAt := time.Now().UnixMicro()
	user.UpdatedAt = &updateAt

	// update user
	_, code, err = u.userRepo.UpdateUser(ctx, id, user)
	if err != nil {
		return nil, code, err
	}

	// get user role
	userRole, code, err := u.userRoleRepo.FindUserRole(ctx, user.RoleUUID, true)
	if err != nil {
		return nil, code, err
	}

	// domain to DTO
	var resp dto.UserResponse
	userDomainToDTORegister(&resp, user, userRole)

	return &resp, http.StatusOK, nil
}
