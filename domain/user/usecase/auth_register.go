package usecase

import (
	"ca-boilerplate/domain"
	"ca-boilerplate/dto"
	passwordhandler "ca-boilerplate/lib/password_handler"
	"context"
	"net/http"
)

func (u *userUsecase) Register(c context.Context, req *dto.UserRegisterRequest) (*dto.UserResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	// check if email already exist
	user, code, err := u.userRepo.FindUserBy(ctx, "email", req.Email, false)
	if err != nil && code != http.StatusNotFound {
		return nil, code, err
	}

	if user != nil {
		return nil, http.StatusConflict, domain.ErrEmailAlreadyExists
	}

	// get user role
	userRole, code, err := u.userRoleRepo.FindUserRoleBy(ctx, "name", "user", false)
	if err != nil {
		return nil, code, err
	}

	// hash password
	pass, err := passwordhandler.HashPassword(req.Password)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	// create user
	var data domain.User
	userDTOtoDomainRegister(&data, req)
	data.Password = pass
	data.RoleUUID = userRole.UUID
	user, code, err = u.userRepo.Register(ctx, &data)
	if err != nil {
		return nil, code, err
	}

	// // create token
	// accesstoken, refreshtoken, err := jwthandler.GenerateAllTokens(user.UUID, userRole.Name)
	// if err != nil {
	// 	return nil, http.StatusInternalServerError, err
	// }

	// // insert RT and AC token to db
	// tokenUUID, code, err := u.tokenRepo.GenerateTokens(ctx, user.UUID, accesstoken, refreshtoken)
	// if err != nil {
	// 	return nil, code, err
	// }

	// _, code, err = u.userRepo.InsertToken(ctx, user.UUID, tokenUUID)
	// if err != nil {
	// 	return nil, code, err
	// }

	// fmt.Println("sini 4")

	// domain to dto
	var resp dto.UserResponse
	userDomainToDTORegister(&resp, user, userRole)

	return &resp, http.StatusCreated, nil
}
