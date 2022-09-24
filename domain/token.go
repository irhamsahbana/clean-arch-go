package domain

import (
	custom "ca-boilerplate/lib/custom_type"
	"context"
)

type Token struct {
	UUID          string               `bson:"uuid"`
	TokenableUUID string               `bson:"tokenable_uuid"`
	TokenableType custom.TokenableType `bson:"tokenable_type"`
	AccessToken   string               `bson:"access_token"`
	RefreshToken  string               `bson:"refresh_token"`
}

type TokenRepositoryContract interface {
	GenerateTokens(ctx context.Context, userId, accessToken, refreshToken string) (uuid string, code int, err error)
	RefreshTokens(ctx context.Context, userId, oldAccessToken, oldRefreshToken, newAccessToken, newRefreshToken string) (uuid string, code int, err error)
	RevokeTokens(ctx context.Context, userId, accessToken string) (uuid string, code int, err error)

	FindTokenWithATandRT(ctx context.Context, accessToken, refreshToken string) (token *Token, code int, err error)
}
