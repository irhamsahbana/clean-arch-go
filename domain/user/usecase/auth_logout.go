package usecase

import (
	"ca-boilerplate/dto"
	"context"
	"errors"
	"net/http"
)

func (u *userUsecase) Logout(c context.Context, userId, accessToken string) (*dto.UserResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	_, code, err := u.userRepo.FindUserBy(ctx, "uuid", userId, false)
	if err != nil {
		return nil, code, err
	}

	if code == http.StatusNotFound {
		return nil, http.StatusUnauthorized, errors.New("Unauthorized")
	}

	tokenUUID, code, err := u.tokenRepo.RevokeTokens(ctx, userId, accessToken)
	if err != nil {
		return nil, code, err
	}

	u.userRepo.RemoveToken(ctx, userId, tokenUUID)

	return nil, code, nil
}
