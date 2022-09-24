package domain

import (
	"ca-boilerplate/dto"
	"context"
	"errors"
)

var ErrEmailAlreadyExists = errors.New("email already exists")

type User struct {
	UUID                   string   `bson:"uuid"`
	RoleUUID               string   `bson:"role_uuid"`
	Name                   string   `bson:"name"`
	Email                  string   `bson:"email"`
	EmailVerifiedAt        *int64   `bson:"email_verified_at"`
	EmailVerificationCodes []string `bson:"email_verification_code"`
	Password               string   `bson:"password"`
	Phone                  *string  `bson:"phone"`
	Whatsapp               *string  `bson:"whatsapp"`
	Tokens                 []string `bson:"tokens"`
	CreatedAt              int64    `bson:"created_at"`
	UpdatedAt              *int64   `bson:"updated_at"`
	DeletedAt              *int64   `bson:"deleted_at"`
}

type UserUsecaseContract interface {
	FindUser(ctx context.Context, id string, withTrashed bool) (*dto.UserResponse, int, error)
	UpdateUser(ctx context.Context, id string, user *dto.UserUpdateRequest) (*dto.UserResponse, int, error)
	DeleteUser(ctx context.Context, id string) (*dto.UserResponse, int, error)

	Register(ctx context.Context, user *dto.UserRegisterRequest) (*dto.UserResponse, int, error)

	Login(ctx context.Context, request *dto.UserLoginRequest) (*dto.UserResponse, int, error)
	RefreshToken(ctx context.Context, oldAccessToken string, oldRefreshToken string, userId string) (*dto.UserResponse, int, error)
	Logout(ctx context.Context, accessToken string, userId string) (*dto.UserResponse, int, error)
}

type UserRepositoryContract interface {
	FindUserBy(ctx context.Context, key string, val interface{}, withTrashed bool) (*User, int, error)
	UpdateUser(ctx context.Context, id string, user *User) (*User, int, error)
	DeleteUser(ctx context.Context, id string) (*User, int, error)

	Register(ctx context.Context, user *User) (*User, int, error)

	InsertToken(ctx context.Context, userId, tokenId string) (*User, int, error)
	RemoveToken(ctx context.Context, userId, tokenId string) (*User, int, error)
}
