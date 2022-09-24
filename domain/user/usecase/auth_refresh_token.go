package usecase

import (
	"ca-boilerplate/dto"
	jwthandler "ca-boilerplate/lib/jwt_handler"
	"context"
	"net/http"
)

func (u *userUsecase) RefreshToken(c context.Context, oldAT, oldRT, userId string) (*dto.UserResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	user, code, err := u.userRepo.FindUserBy(ctx, "uuid", userId, false)
	if err != nil {
		return nil, code, err
	}

	userRole, code, err := u.userRoleRepo.FindUserRole(ctx, user.RoleUUID, false)
	if err != nil {
		return nil, code, err
	}

	token, code, err := u.tokenRepo.FindTokenWithATandRT(ctx, oldAT, oldRT)
	if err != nil {
		return nil, code, err
	}

	accesstoken, refreshtoken, err := jwthandler.GenerateAllTokens(user.UUID, userRole.Name)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	tokenUUID, code, err := u.tokenRepo.RefreshTokens(ctx, user.UUID, oldAT, oldRT, accesstoken, refreshtoken)
	if err != nil {
		return nil, code, err
	}

	_, code, err = u.userRepo.RemoveToken(ctx, user.UUID, token.UUID)
	if err != nil {
		return nil, code, err
	}

	_, code, err = u.userRepo.InsertToken(ctx, user.UUID, tokenUUID)
	if err != nil {
		return nil, code, err
	}

	var resp dto.UserResponse
	userDomainToDTOLogin(&resp, user, userRole, &accesstoken, &refreshtoken)

	return &resp, http.StatusOK, nil
}
